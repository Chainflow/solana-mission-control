package exporter

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	// "k8s.io/klog/v2"

	"github.com/Chainflow/solana-mission-control/alerter"
	"github.com/Chainflow/solana-mission-control/config"
	"github.com/Chainflow/solana-mission-control/monitor"
	"github.com/Chainflow/solana-mission-control/utils"
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
		Help: "Current epoch of validator (max confirmation)",
	})

	networkEpoch = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_network_epoch",
		Help: "Current epoch of network (max confirmation)",
	})

	epochDifference = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_epoch_diff",
		Help: "Current epoch difference of network and validator (max confirmation)",
	})

	epochFirstSlot = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_confirmed_epoch_first_slot",
		Help: "Current epoch's first slot (max confirmation) - validator",
	})

	epochLastSlot = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_confirmed_epoch_last_slot",
		Help: "Current epoch's last slot (max confirmation) - validator",
	})

	networkEpochLastSlot = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_network_confirmed_epoch_last_slot",
		Help: "Confirmed epoch's last slot (max confirmation) - network",
	})

	nodeHealth = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_node_health",
		Help: "Current health of the node.",
	})

	balance = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "account_balance",
		Help: "Current balance of your account.",
	})

	leaderSlotsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "solana_leader_slots_total",
			Help: "Number of leader slots per leader, grouped by skip status (max confirmation)",
		},
		[]string{"status", "nodekey"})

	valBlockHeight = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_block_height",
		Help: "Current Block Height of validator",
	})

	networkBlockHeight = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_network_block_height",
		Help: "Current Block Height of network",
	})

	blockDiff = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_block_height_diff",
		Help: "Current Block Height difference of network and validator",
	})

	valSkipRate = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_val_skip_rate",
		Help: "Validator skip rate",
	})

	netSkipRate = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_net_skip_rate",
		Help: "Network skip rate",
	})

	skipRateDifference = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_skip_rate_diff",
		Help: "Skip rate difference of network and validator",
	})

	leaderSlots = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_val_leader_slots",
		Help: "Leader slots of a validator in current epoch",
	})

	totalSlots = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_total_slots",
		Help: "Total slots in current epoch",
	})

	valBlocksProduced = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_val_blocks_produced",
		Help: "Blocks produced of a validator in current epoch",
	})

	totalBlocksProduced = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_total_blocks_produced",
		Help: "Total blocks produced in current epoch",
	})

	skippdSlots = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_val_skipped_slots",
		Help: "Skipped slots of a validator in current epoch",
	})

	skippedTotal = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "solana_skipped_total",
		Help: "Total skipped slots in current epoch",
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
	prometheus.MustRegister(valBlockHeight)
	prometheus.MustRegister(networkBlockHeight)
	prometheus.MustRegister(networkEpoch)
	prometheus.MustRegister(epochDifference)
	prometheus.MustRegister(blockDiff)
	prometheus.MustRegister(valSkipRate)
	prometheus.MustRegister(netSkipRate)
	prometheus.MustRegister(skipRateDifference)
	prometheus.MustRegister(leaderSlots)
	prometheus.MustRegister(totalSlots)
	prometheus.MustRegister(valBlocksProduced)
	prometheus.MustRegister(totalBlocksProduced)
	prometheus.MustRegister(skippdSlots)
	prometheus.MustRegister(skippedTotal)
	prometheus.MustRegister(networkEpochLastSlot)
}

// WatchSlots get data from different methods and store that data in prometheus. Those are
// 1. Account balance
// 2. Node Health
// 3. Network Epoch Information
// 4. Validator Epoch information
// 5. epoch difference of network and validator and send alert if it is drppoed below epoch threshold
// 6. block height difference of network and validator
// 7. fetch a new leader schedule if previous epoch has done
// 8. Get list of confirmed blocks
// 9. Skip Rate of the validator and network and difference between them
// 10. Leader slots and total slots of the validator
// 11. Valid and total blocks produced
// 12. Skipped slots and total slots skipped
func (c *solanaCollector) WatchSlots(cfg *config.Config) {
	ticker := time.NewTicker(slotPacerSchedule)

	for {
		<-ticker.C

		// Get identity account balance
		bal, err := monitor.GetIdentityBalance(cfg)
		if err != nil {
			log.Printf("Error while getting account balance : %v", err)
			continue
		}

		balance.Set(float64(bal.Result.Value) / math.Pow(10, 9))

		// Get skip rate of validator and network using solana cli command
		valSkip, netSkip, err := monitor.SkipRate(cfg)
		if err != nil {
			log.Printf("Error while getting skipped slots : %v", err)
			continue
		}
		valSkipRate.Set(valSkip)
		netSkipRate.Set(netSkip)
		skipdiff := valSkip - netSkip
		// if skipdiff > 0 {
		// 	skipRateDifference.Set(skipdiff)
		// } else {
		skipRateDifference.Set(skipdiff)
		// }

		// Get Node Health
		h, err := monitor.GetNodeHealth(cfg)
		if err != nil {
			log.Printf("Error while getting node health info : %v", err)
			continue
		}

		nodeHealth.Set(h) // set node health

		// Get network epoch info
		resp, err := monitor.GetEpochInfo(cfg, utils.Network)
		if err != nil {
			log.Printf("failed to fetch epoch info of network, retrying: %v", err)
			// cancel()
			continue
		}

		networkEpoch.Set(float64(resp.Result.Epoch))             // Set n/w epoch
		networkBlockHeight.Set(float64(resp.Result.BlockHeight)) // set n/w block height

		// Calculate first and last slot in network epoch.
		netFirstSlot := resp.Result.AbsoluteSlot - resp.Result.SlotIndex
		netLastSlot := netFirstSlot + resp.Result.SlotsInEpoch
		networkEpochLastSlot.Set(float64(netLastSlot)) // set confirmed epoch last slock - network

		// Get recent block production details
		bp, err := monitor.BlockProduction(cfg)
		if err != nil {
			log.Printf("Error while getting block production details : %v", err)
		}

		leaderSlots.Set(float64(bp.LeaderSlots))
		totalSlots.Set(float64(bp.TotalSlots))
		valBlocksProduced.Set(float64(bp.BlocksProduced))
		totalBlocksProduced.Set(float64(bp.TotalBlocksProduced))
		skippdSlots.Set(float64(bp.SkippedSlots))
		skippedTotal.Set(float64(bp.TotalSlotsSkipped))

		// Get validator epoch info
		resp, err = monitor.GetEpochInfo(cfg, utils.Validator)
		if err != nil {
			log.Printf("failed to fetch poch info of validator, retrying: %v", err)
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
		valBlockHeight.Set(float64(info.BlockHeight))

		log.Printf("************** Block Height ********* : %d", info.BlockHeight)

		// Calculate epoch difference of network and validator
		diff := float64(resp.Result.Epoch) - float64(info.Epoch)
		epochDifference.Set(diff) // set epoch diff to prometheus

		if strings.EqualFold(cfg.AlerterPreferences.EpochDiffAlerts, "yes") && int64(diff) >= cfg.AlertingThresholds.EpochDiffThreshold && int64(diff) > 0 {
			// send alert
			err = alerter.SendTelegramAlert(fmt.Sprintf("Epoch Difference Alert : Difference b/w network and validator epoch has exceeded the configured thershold %d", cfg.AlertingThresholds.EpochDiffThreshold), cfg)
			if err != nil {
				log.Printf("Error while sending epoch diff alert to telegram: %v", err)
			}
			// send email alert
			err = alerter.SendEmailAlert(fmt.Sprintf("Epoch Difference Alert : Difference b/w network and validator epoch has exceeded the configured thershold %d", cfg.AlertingThresholds.EpochDiffThreshold), cfg)
			if err != nil {
				log.Printf("Error while sending epoch diff alert to email: %v", err)
			}
		}

		heightDiff := float64(resp.Result.BlockHeight) - float64(info.BlockHeight)
		blockDiff.Set(heightDiff) // block height difference of network and validator

		if int64(heightDiff) >= cfg.AlertingThresholds.BlockDiffThreshold {
			// send alert
			err = alerter.SendTelegramAlert(fmt.Sprintf("Block Difference Alert : Block difference b/w network and validator has exceeded %d", cfg.AlertingThresholds.BlockDiffThreshold), cfg)
			if err != nil {
				log.Printf("Error while sending block height diff alert to telegram: %v", err)
			}

			// send email alert
			err = alerter.SendEmailAlert(fmt.Sprintf("Block Difference Alert : Block difference b/w network and validator has exceeded %d", cfg.AlertingThresholds.BlockDiffThreshold), cfg)
			if err != nil {
				log.Printf("Error while sending block height diff alert to email: %v", err)
			}
		}
	}
}
