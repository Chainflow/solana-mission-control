package exporter

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	// "k8s.io/klog/v2"

	"github.com/PrathyushaLakkireddy/solana-prometheus/alerter"
	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/monitor"
	"github.com/PrathyushaLakkireddy/solana-prometheus/types"
)

const (
	slotPacerSchedule = 5 * time.Second
)

var (
	confirmedSlotHeight = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_confirmed_slot_height",
		Help: "Last confirmed slot height processed by watcher routine (max confirmation)",
	})

	currentEpochNumber = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_confirmed_epoch_number",
		Help: "Current epoch (max confirmation)",
	})

	epochFirstSlot = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_confirmed_epoch_first_slot",
		Help: "Current epoch's first slot (max confirmation)",
	})

	epochLastSlot = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_confirmed_epoch_last_slot",
		Help: "Current epoch's last slot (max confirmation)",
	})

	nodeHealth = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_node_health",
		Help: "Current health of the node.",
	})

	balance = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "Account_balance",
		Help: "Current balance if your account.",
	})

	leaderSlotsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "solana_leader_slots_total",
			Help: "Number of leader slots per leader, grouped by skip status (max confirmation)",
		},
		[]string{"status", "nodekey"})

	blockHeight = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_block_height",
		Help: "Current Block Height.",
	})
)

func init() {
	prometheus.MustRegister(confirmedSlotHeight)
	prometheus.MustRegister(currentEpochNumber)
	prometheus.MustRegister(epochFirstSlot)
	prometheus.MustRegister(epochLastSlot)
	prometheus.MustRegister(leaderSlotsTotal)
	prometheus.MustRegister(nodeHealth)
	prometheus.MustRegister(balance)
	prometheus.MustRegister(blockHeight)
}

func (c *solanaCollector) WatchSlots(cfg *config.Config) {
	var (
		// Current mapping of relative slot numbers to leader public keys.
		epochSlots map[int64]string
		// Current epoch number corresponding to epochSlots.
		epochNumber int64
		// Last slot number we generated ticks for.
		watermark int64
	)

	ticker := time.NewTicker(slotPacerSchedule)

	for {
		<-ticker.C

		// Get account balance
		bal, err := monitor.GetBalance(cfg)
		if err != nil {
			log.Printf("Error while getting account balance : %v", err)
			continue
		}

		balance.Set(float64(bal.Result.Value))

		// Get Node Health
		health, err := monitor.GetNodeHealth(cfg)
		if err != nil {
			log.Printf("Error while getting node health info : %v", err)
			continue
		}
		var h float64
		if health.Error.Message != "" {
			if strings.EqualFold(health.Error.Message, "Node is unhealthy") {
				h = 0
			}
		} else {
			if strings.EqualFold(health.Result, "ok") {
				h = 1
			}
		}

		nodeHealth.Set(h)

		resp, err := monitor.GetEpochInfo(cfg)
		if err != nil {
			log.Printf("failed to fetch info info, retrying: %v", err)
			// cancel()
			continue
		}
		// cancel()
		info := resp.Result

		// Calculate first and last slot in epoch.
		firstSlot := info.AbsoluteSlot - info.SlotIndex
		lastSlot := firstSlot + info.SlotsInEpoch

		confirmedSlotHeight.Set(float64(info.AbsoluteSlot))
		currentEpochNumber.Set(float64(info.Epoch))
		epochFirstSlot.Set(float64(firstSlot))
		epochLastSlot.Set(float64(lastSlot))
		blockHeight.Set(float64(resp.Result.BlockHeight))

		// Check whether we need to fetch a new leader schedule
		if epochNumber != info.Epoch {
			log.Printf("new epoch at slot %d: %d (previous: %d)", firstSlot, info.Epoch, epochNumber)

			epochSlots, err := monitor.GetLeaderSlots(firstSlot, cfg)
			if err != nil {
				log.Printf("failed to request leader schedule, retrying: %v", err)
				continue
			}

			log.Printf("%d leader slots in epoch %d", len(epochSlots), info.Epoch)

			epochNumber = info.Epoch
			log.Printf("we're still in epoch %d, not fetching leader schedule", info.Epoch)

			// Reset watermark to current offset on new epoch (we do not backfill slots we missed at startup)
			watermark = info.SlotIndex
		} else if watermark == info.SlotIndex {
			log.Printf("slot has not advanced at %d, skipping", info.AbsoluteSlot)
			continue
		}

		log.Printf("confirmed slot %d (offset %d, +%d), epoch %d (from slot %d to %d, %d remaining)",
			info.AbsoluteSlot, info.SlotIndex, info.SlotIndex-watermark, info.Epoch, firstSlot, lastSlot, lastSlot-info.AbsoluteSlot)

		// Get list of confirmed blocks since the last request. This is totally undocumented, but the result won't
		// contain missed blocks, allowing us to figure out block production success rate.
		rangeStart := firstSlot + watermark
		rangeEnd := firstSlot + info.SlotIndex - 1

		cfm, err := monitor.GetConfirmedBlocks(rangeStart, rangeEnd, cfg)
		if err != nil {
			log.Printf("failed to request confirmed blocks at %d, retrying: %v", watermark, err)
			// cancel()
			continue
		}

		log.Printf("confirmed blocks: %d -> %d: %v", rangeStart, rangeEnd, cfm)

		// Figure out leaders for each block in range
		for i := watermark; i < info.SlotIndex; i++ {
			leader, ok := epochSlots[i]
			abs := firstSlot + i
			if !ok {
				// This cannot happen with a well-behaved node and is a programming error in either Solana or the exporter.
				log.Printf("slot %d (offset %d) missing from epoch %d leader schedule",
					abs, i, info.Epoch)
			}

			// Check if block was included in getConfirmedBlocks output, otherwise, it was skipped.
			var present bool
			for _, s := range cfm {
				if abs == s {
					present = true
				}
			}

			var skipped string
			var label string
			if present {
				skipped = "(valid)"
				label = "valid"
			} else {
				skipped = "(SKIPPED)"
				label = "skipped"
			}

			leaderSlotsTotal.With(prometheus.Labels{"status": label, "nodekey": leader}).Add(1)
			log.Printf("slot %d (offset %d) with leader %s %s", abs, i, leader, skipped)
		}

		watermark = info.SlotIndex
	}
}
