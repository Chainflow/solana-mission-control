package exporter

import (
	// "context"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	// "github.com/certusone/solana_exporter/pkg/rpc"
	"github.com/prometheus/client_golang/prometheus"
	// "k8s.io/klog/v2"

	"github.com/PrathyushaLakkireddy/solana-prometheus/alerter"
	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/monitor"
	"github.com/PrathyushaLakkireddy/solana-prometheus/types"
	"github.com/PrathyushaLakkireddy/solana-prometheus/utils"
)

const (
	httpTimeout = 5 * time.Second
)

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
	commission              *prometheus.Desc
	delinqentCommission     *prometheus.Desc
	validatorVote           *prometheus.Desc
	StatusAlertCount        *prometheus.Desc
	txCount                 *prometheus.Desc
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
			[]string{"votekey", "pubkey"}, nil),
		validatorLastVote: prometheus.NewDesc(
			"solana_validator_last_vote",
			"Last voted slot per validator",
			[]string{"votekey", "pubkey"}, nil),
		validatorRootSlot: prometheus.NewDesc(
			"solana_validator_root_slot",
			"Root slot per validator",
			[]string{"votekey", "pubkey"}, nil),
		validatorDelinquent: prometheus.NewDesc(
			"solana_validator_delinquent",
			"Whether a validator is delinquent",
			[]string{"votekey", "pubkey"}, nil),
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
		commission: prometheus.NewDesc(
			"solana_val_commission",
			"Solana validator current commission.",
			[]string{"solana_val_commission"}, nil,
		),
		delinqentCommission: prometheus.NewDesc(
			"solana_val_delinquuent_commission",
			"Solana validator delinqent commission.",
			[]string{"solana_delinquent_commission"}, nil,
		),
		validatorVote: prometheus.NewDesc(
			"solana_vote_account",
			"whether the vote account is staked for this epoch",
			[]string{"state"}, nil,
		),
		StatusAlertCount: prometheus.NewDesc(
			"solana_val_alert_count",
			"Count of alerts about validator status alerting",
			[]string{"alert_count"}, nil,
		),
		txCount: prometheus.NewDesc(
			"solana_tx_count",
			"Current transaction count from the ledger.",
			[]string{"solana_tx_count"}, nil,
		),
	}
}

func (c *solanaCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.solanaVersion
	ch <- c.accountBalance
	ch <- c.totalValidatorsDesc
	ch <- c.slotLeader
	ch <- c.currentSlot
	ch <- c.commission
	ch <- c.delinqentCommission
	ch <- c.validatorVote
	// ch <- c.StatusAlertCount
	ch <- c.txCount
}

func (c *solanaCollector) mustEmitMetrics(ch chan<- prometheus.Metric, response types.GetVoteAccountsResponse) {
	ch <- prometheus.MustNewConstMetric(c.totalValidatorsDesc, prometheus.GaugeValue,
		float64(len(response.Result.Delinquent)), "delinquent")
	ch <- prometheus.MustNewConstMetric(c.totalValidatorsDesc, prometheus.GaugeValue,
		float64(len(response.Result.Current)), "current")

	count, _ := monitor.GetTxCount(c.config)
	txcount := utils.NearestThousandFormat(float64(count.Result))

	ch <- prometheus.MustNewConstMetric(c.txCount, prometheus.GaugeValue, float64(count.Result), txcount)

	for _, account := range append(response.Result.Current, response.Result.Delinquent...) {
		if account.NodePubkey == c.config.ValDetails.PubKey {
			// ch <- prometheus.MustNewConstMetric(c.validatorActivatedStake, prometheus.GaugeValue,
			// 	float64(account.ActivatedStake), account.VotePubkey, account.NodePubkey)
			ch <- prometheus.MustNewConstMetric(c.validatorLastVote, prometheus.GaugeValue,
				float64(account.LastVote), account.VotePubkey, account.NodePubkey)
			ch <- prometheus.MustNewConstMetric(c.validatorRootSlot, prometheus.GaugeValue,
				float64(account.RootSlot), account.VotePubkey, account.NodePubkey)
		}
	}
	var epochvote float64
	for _, vote := range response.Result.Current {
		if vote.NodePubkey == c.config.ValDetails.PubKey {
			v := strconv.FormatInt(vote.Commission, 10)

			if vote.EpochVoteAccount {
				epochvote = 1
			} else {
				epochvote = 0
			}
			ch <- prometheus.MustNewConstMetric(c.validatorVote, prometheus.GaugeValue,
				epochvote, "current")
			ch <- prometheus.MustNewConstMetric(c.commission, prometheus.GaugeValue, float64(vote.Commission), v)

			ch <- prometheus.MustNewConstMetric(c.validatorDelinquent, prometheus.GaugeValue,
				0, vote.VotePubkey, vote.NodePubkey)

			stake := float64(vote.ActivatedStake) / math.Pow(10, 9)

			ch <- prometheus.MustNewConstMetric(c.validatorActivatedStake, prometheus.GaugeValue,
				stake, vote.VotePubkey, vote.NodePubkey)

			// Check weather the validator is voting or not
			time.NewTimer(180 * time.Second)
			if vote.EpochVoteAccount == false && vote.ActivatedStake <= 0 {
				//time.NewTimer(180 * time.Second)
				msg := "Solana validator is NOT VOTING"
				c.AlertValidatorStatus(msg, ch)
			} else {
				//time.NewTimer(180 * time.Second)
				msg := "Solana validator is VOTING"
				c.AlertValidatorStatus(msg, ch)

			}

		}
	}

	for _, vote := range response.Result.Delinquent {
		if vote.NodePubkey == c.config.ValDetails.PubKey {
			v := strconv.FormatInt(vote.Commission, 10)
			if vote.EpochVoteAccount {
				epochvote = 1
			} else {
				epochvote = 0
			}
			ch <- prometheus.MustNewConstMetric(c.validatorVote, prometheus.GaugeValue,
				epochvote, "delinquent")
			ch <- prometheus.MustNewConstMetric(c.delinqentCommission, prometheus.GaugeValue, float64(vote.Commission), v)

			ch <- prometheus.MustNewConstMetric(c.validatorDelinquent, prometheus.GaugeValue,
				1, vote.VotePubkey, vote.NodePubkey)

<<<<<<< HEAD
			err := alerter.SendTelegramAlert(fmt.Sprintf("Your solana validator is in DELINQUENT state"), c.config)
			if err != nil {
				log.Printf("Error while sending vallidator status alert: %v", err)
			}
=======
>>>>>>> validator alerts modified
		}
	}
}

func (c *solanaCollector) AlertValidatorStatus(msg string, ch chan<- prometheus.Metric) {
	now := time.Now().UTC()
	currentTime := now.Format(time.Kitchen)

	var alertsArray []string

	for _, value := range c.config.RegularStatusAlerts.AlertTimings {
		t, _ := time.Parse(time.Kitchen, value)
		alertTime := t.Format(time.Kitchen)

		alertsArray = append(alertsArray, alertTime)
	}

	log.Printf("Current time : %v and alerts array : %v", currentTime, alertsArray)

	var count float64 = 0

	for _, statusAlertTime := range alertsArray {
		if currentTime == statusAlertTime {
<<<<<<< HEAD
			count = count + 1

			// log.Println("count....", count)

			// ch <- prometheus.MustNewConstMetric(c.StatusAlertCount, prometheus.GaugeValue,
			// 	count, "1")

			count1, _ := monitor.AlertStatusCountFromPrometheus(c.config)
			if count1 == "1" {
				return
			}

			log.Printf("count1..", count1)

			err := alerter.SendTelegramAlert(msg, c.config)
			if err != nil {
				log.Printf("Error while sending vallidator status alert: %v", err)
			}

		} else {
			// ch <- prometheus.MustNewConstMetric(c.StatusAlertCount, prometheus.GaugeValue,
			// 	count, "0")
=======
			dbcount, _ := monitor.AlertStatusCountFromPrometheus(c.config)
			if dbcount == "false" {
				err := alerter.SendTelegramAlert(msg, c.config)
				if err != nil {
					log.Printf("Error while sending vallidator status alert: %v", err)
				}
				ch <- prometheus.MustNewConstMetric(c.StatusAlertCount, prometheus.GaugeValue,
					count, "true")
				count = count + 1
			} else {
				ch <- prometheus.MustNewConstMetric(c.StatusAlertCount, prometheus.GaugeValue,
					count, "false")
				return
			}
>>>>>>> validator alerts modified
		}
		// else {
		// 	ch <- prometheus.MustNewConstMetric(c.StatusAlertCount, prometheus.GaugeValue,
		// 		count, "0")
		// }
	}
}

func (c *solanaCollector) Collect(ch chan<- prometheus.Metric) {
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
		log.Printf("Error while getting block time: %v", err)
	}

	pvt, err := monitor.GetBlockTime(slot.Result-1, c.config)
	if err != nil {
		log.Printf("Error while getting previous block time: %v", err)
	}

	t1 := time.Unix(bt.Result, 0)

	t2 := time.Unix(pvt.Result, 0)

	sub := t1.Sub(t2)

	diff := sub.Seconds()

	fmt.Println("tdiff  time is .....", diff)

	if diff < 0 {
		diff = -(diff)
	}
	s := fmt.Sprintf("%.2f", diff)
	sec, _ := strconv.ParseFloat(s, 64)
	ch <- prometheus.MustNewConstMetric(c.blockTime, prometheus.GaugeValue, sec, s+"s")
}
