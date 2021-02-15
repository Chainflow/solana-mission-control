package exporter

import (
	// "context"
	"fmt"
	"log"
	"strconv"
	"time"

	// "github.com/certusone/solana_exporter/pkg/rpc"
	"github.com/prometheus/client_golang/prometheus"
	// "k8s.io/klog/v2"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/monitor"
	"github.com/PrathyushaLakkireddy/solana-prometheus/types"
)

const (
	httpTimeout = 5 * time.Second
)

// var (
// 	rpcAddr = cfg.
// 	addr    = flag.String("addr", ":8080", "Listen address")
// )

// func init() {
// 	klog.InitFlags(nil)
// }

type solanaCollector struct {
	config                  *config.Config
	totalValidatorsDesc     *prometheus.Desc
	validatorActivatedStake *prometheus.Desc
	validatorLastVote       *prometheus.Desc
	validatorRootSlot       *prometheus.Desc
	validatorDelinquent     *prometheus.Desc
	solanaVersion           *prometheus.Desc
	accountBalance          *prometheus.Desc
	slotLeader              *prometheus.Desc
	blockTime               *prometheus.Desc
	currentSlot             *prometheus.Desc
}

func NewSolanaCollector(cfg *config.Config) *solanaCollector {
	return &solanaCollector{
		config: cfg,
		totalValidatorsDesc: prometheus.NewDesc(
			"solana_active_validators",
			"Total number of active validators by state",
			[]string{"state"}, nil),
		validatorActivatedStake: prometheus.NewDesc(
			"solana_validator_activated_stake",
			"Activated stake per validator",
			[]string{"pubkey", "nodekey"}, nil),
		validatorLastVote: prometheus.NewDesc(
			"solana_validator_last_vote",
			"Last voted slot per validator",
			[]string{"pubkey", "nodekey"}, nil),
		validatorRootSlot: prometheus.NewDesc(
			"solana_validator_root_slot",
			"Root slot per validator",
			[]string{"pubkey", "nodekey"}, nil),
		validatorDelinquent: prometheus.NewDesc(
			"solana_validator_delinquent",
			"Whether a validator is delinquent",
			[]string{"pubkey", "nodekey"}, nil),
		solanaVersion: prometheus.NewDesc(
			"solana_node_version",
			"Node version of solana",
			[]string{"version"}, nil),
		accountBalance: prometheus.NewDesc(
			"solana_account_balance",
			"Account balance",
			[]string{"solana_acc_balance"}, nil),
		slotLeader: prometheus.NewDesc(
			"solana_slot_leader",
			"Current slot leader",
			[]string{"solana_slot_leader"}, nil),
		currentSlot: prometheus.NewDesc(
			"solana_current_slot",
			"Current slot height",
			[]string{"solana_current_slot"}, nil,
		),
		blockTime: prometheus.NewDesc(
			"solana_block_time",
			"Current block time.",
			[]string{"solana_block_time"}, nil,
		),
	}
}

func (c *solanaCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.solanaVersion
	ch <- c.accountBalance
	ch <- c.totalValidatorsDesc
	ch <- c.slotLeader
	ch <- c.currentSlot
}

func (c *solanaCollector) mustEmitMetrics(ch chan<- prometheus.Metric, response types.GetVoteAccountsResponse) {
	ch <- prometheus.MustNewConstMetric(c.totalValidatorsDesc, prometheus.GaugeValue,
		float64(len(response.Result.Delinquent)), "delinquent")
	ch <- prometheus.MustNewConstMetric(c.totalValidatorsDesc, prometheus.GaugeValue,
		float64(len(response.Result.Current)), "current")

	for _, account := range append(response.Result.Current, response.Result.Delinquent...) {
		if account.VotePubkey == c.config.ValDetails.PubKey {
			ch <- prometheus.MustNewConstMetric(c.validatorActivatedStake, prometheus.GaugeValue,
				float64(account.ActivatedStake), account.VotePubkey, account.NodePubkey)
			ch <- prometheus.MustNewConstMetric(c.validatorLastVote, prometheus.GaugeValue,
				float64(account.LastVote), account.VotePubkey, account.NodePubkey)
			ch <- prometheus.MustNewConstMetric(c.validatorRootSlot, prometheus.GaugeValue,
				float64(account.RootSlot), account.VotePubkey, account.NodePubkey)
		}
	}
	for _, account := range response.Result.Current {
		if account.VotePubkey == c.config.ValDetails.PubKey {
			ch <- prometheus.MustNewConstMetric(c.validatorDelinquent, prometheus.GaugeValue,
				0, account.VotePubkey, account.NodePubkey)
		}
	}
	for _, account := range response.Result.Delinquent {
		if account.VotePubkey == c.config.ValDetails.PubKey {
			ch <- prometheus.MustNewConstMetric(c.validatorDelinquent, prometheus.GaugeValue,
				1, account.VotePubkey, account.NodePubkey)
		}
	}
}

func (c *solanaCollector) Collect(ch chan<- prometheus.Metric) {
	// ctx, cancel := context.WithTimeout(context.Background(), httpTimeout)
	// defer cancel()

	accs, err := monitor.GetVoteAccounts(c.config)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.totalValidatorsDesc, err)
		ch <- prometheus.NewInvalidMetric(c.validatorActivatedStake, err)
		ch <- prometheus.NewInvalidMetric(c.validatorLastVote, err)
		ch <- prometheus.NewInvalidMetric(c.validatorRootSlot, err)
		ch <- prometheus.NewInvalidMetric(c.validatorDelinquent, err)
	} else {
		c.mustEmitMetrics(ch, accs)
	}

	version, err := monitor.GetVersion(c.config)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.solanaVersion, err)
	} else {
		ch <- prometheus.MustNewConstMetric(c.solanaVersion, prometheus.GaugeValue, 1, version.Result.SolanaCore)
	}

	bal, err := monitor.GetBalance(c.config)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.accountBalance, err)
	} else {
		s := strconv.FormatInt(bal.Result.Value, 10)
		ch <- prometheus.MustNewConstMetric(c.accountBalance, prometheus.GaugeValue, 1, s)
	}

	leader, err := monitor.GetSlotLeader(c.config)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.slotLeader, err)
	} else {
		ch <- prometheus.MustNewConstMetric(c.slotLeader, prometheus.GaugeValue, 1, leader.Result)
	}

	slot, err := monitor.GetCurrentSlot(c.config)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.currentSlot, err)
	} else {
		cs := strconv.FormatInt(slot.Result, 10)
		ch <- prometheus.MustNewConstMetric(c.currentSlot, prometheus.GaugeValue, 1, cs)
	}

	bt, err := monitor.GetBlockTime(slot.Result, c.config)
	if err != nil {
		log.Printf("Errir while getting block time: %v", err)
	}

	pvt, err := monitor.GetBlockTime(slot.Result-1, c.config)
	if err != nil {
		log.Printf("Errir while getting previous block time: %v", err)
	}

	t1 := time.Unix(bt.Result, 0)
	t2 := time.Unix(pvt.Result, 0)

	sub := t1.Sub(t2)
	diff := sub.Seconds()

	if diff < 0 {
		diff = -(diff)
	}
	s := fmt.Sprintf("%.2f", diff)

	ch <- prometheus.MustNewConstMetric(c.blockTime, prometheus.GaugeValue, diff, s+"s")
}
