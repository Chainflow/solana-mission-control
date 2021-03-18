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
	statusAlertCount        *prometheus.Desc
	ipAddress               *prometheus.Desc
	txCount                 *prometheus.Desc
	netVoteHeight           *prometheus.Desc
	valVoteHeight           *prometheus.Desc
	voteHeightDiff          *prometheus.Desc
	valVotingStatus         *prometheus.Desc
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
		statusAlertCount: prometheus.NewDesc(
			"solana_val_alert_count",
			"Count of alerts about validator status alerting",
			[]string{"alert_count"}, nil,
		),
		ipAddress: prometheus.NewDesc(
			"solana_ip_address",
			"IP Address from clustrnode information, gossip",
			[]string{"ip_address"}, nil,
		),
		txCount: prometheus.NewDesc(
			"solana_tx_count",
			"solana transaction count",
			[]string{"solana_tx_count"}, nil,
		),
		netVoteHeight: prometheus.NewDesc(
			"solana_network_vote_height",
			"solana network vote height",
			[]string{"solana_network_vote_height"}, nil,
		),
		valVoteHeight: prometheus.NewDesc(
			"solana_validator_vote_height",
			"solana validator vote height",
			[]string{"solana_validator_vote_height"}, nil,
		),
		voteHeightDiff: prometheus.NewDesc(
			"solana_vote_height_diff",
			"solana vote height difference of validator and network",
			[]string{"solana_vote_height_diff"}, nil,
		),
		valVotingStatus: prometheus.NewDesc(
			"solana_val_status",
			"Validator voting status i.e., voting or jailed.",
			[]string{"solana_val_status"}, nil,
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
	ch <- c.ipAddress
	// ch <- c.StatusAlertCount
	ch <- c.txCount
	ch <- c.netVoteHeight
	ch <- c.valVoteHeight
	ch <- c.voteHeightDiff
	ch <- c.valVotingStatus
}

func (c *solanaCollector) mustEmitMetrics(ch chan<- prometheus.Metric, response types.GetVoteAccountsResponse) {
	ch <- prometheus.MustNewConstMetric(c.totalValidatorsDesc, prometheus.GaugeValue,
		float64(len(response.Result.Delinquent)), "delinquent")
	ch <- prometheus.MustNewConstMetric(c.totalValidatorsDesc, prometheus.GaugeValue,
		float64(len(response.Result.Current)), "current")

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
	var valresult float64
	// current vote account information
	for _, vote := range response.Result.Current {
		if vote.NodePubkey == c.config.ValDetails.PubKey {
			v := strconv.FormatInt(vote.Commission, 10)

			if vote.EpochVoteAccount {
				epochvote = 1
			} else {
				epochvote = 0
			}
			ch <- prometheus.MustNewConstMetric(c.validatorVote, prometheus.GaugeValue,
				epochvote, "current") // store vote account is staked or not

			ch <- prometheus.MustNewConstMetric(c.commission, prometheus.GaugeValue, float64(vote.Commission), v) // store commission

			ch <- prometheus.MustNewConstMetric(c.validatorDelinquent, prometheus.GaugeValue,
				0, vote.VotePubkey, vote.NodePubkey) // stor vote key and node key

			stake := float64(vote.ActivatedStake) / math.Pow(10, 9)
			ch <- prometheus.MustNewConstMetric(c.validatorActivatedStake, prometheus.GaugeValue,
				stake, vote.VotePubkey, vote.NodePubkey) // store activated stake

			// Check weather the validator is voting or not
			if vote.EpochVoteAccount == false && vote.ActivatedStake <= 0 {
				msg := "Solana validator is NOT VOTING"
				c.AlertValidatorStatus(msg, ch)

				ch <- prometheus.MustNewConstMetric(c.valVotingStatus, prometheus.GaugeValue, 0, "Jailed")
			} else {
				msg := "Solana validator is VOTING"
				c.AlertValidatorStatus(msg, ch)

				ch <- prometheus.MustNewConstMetric(c.valVotingStatus, prometheus.GaugeValue, 1, "Voting")
			}
			valresult = float64(vote.LastVote)
			ch <- prometheus.MustNewConstMetric(c.valVoteHeight, prometheus.GaugeValue, valresult, "validator")
			netresult := c.getNetworkVoteAccountinfo()
			ch <- prometheus.MustNewConstMetric(c.netVoteHeight, prometheus.GaugeValue, netresult, "network")
			diff := netresult - valresult
			ch <- prometheus.MustNewConstMetric(c.voteHeightDiff, prometheus.GaugeValue, diff, "vote height difference")

		}
	}

	// delinquent vote account information
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
			ch <- prometheus.MustNewConstMetric(c.delinqentCommission, prometheus.GaugeValue, float64(vote.Commission), v) // store delinquent commission

			// send alert if the validator is delinquent
			ch <- prometheus.MustNewConstMetric(c.validatorDelinquent, prometheus.GaugeValue,
				1, vote.VotePubkey, vote.NodePubkey)

			err := alerter.SendTelegramAlert(fmt.Sprintf("Your solana validator is in DELINQUENT state"), c.config)
			if err != nil {
				log.Printf("Error while sending vallidator status alert: %v", err)
			}

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
			dbcount, _ := monitor.AlertStatusCountFromPrometheus(c.config)
			if dbcount == "false" {
				err := alerter.SendTelegramAlert(msg, c.config)
				if err != nil {
					log.Printf("Error while sending vallidator status alert: %v", err)
				}
				ch <- prometheus.MustNewConstMetric(c.statusAlertCount, prometheus.GaugeValue,
					count, "true")
				count = count + 1
			} else {
				ch <- prometheus.MustNewConstMetric(c.statusAlertCount, prometheus.GaugeValue,
					count, "false")
				return
			}
		}
		// else {
		// 	ch <- prometheus.MustNewConstMetric(c.StatusAlertCount, prometheus.GaugeValue,
		// 		count, "0")
		// }
	}
}

func (c *solanaCollector) Collect(ch chan<- prometheus.Metric) {
	accs, err := monitor.GetVoteAccounts(c.config, utils.Validator) // get vote accounts
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.totalValidatorsDesc, err)
		ch <- prometheus.NewInvalidMetric(c.validatorActivatedStake, err)
		ch <- prometheus.NewInvalidMetric(c.validatorLastVote, err)
		ch <- prometheus.NewInvalidMetric(c.validatorRootSlot, err)
		ch <- prometheus.NewInvalidMetric(c.validatorDelinquent, err)
	} else {
		c.mustEmitMetrics(ch, accs) // emit vote account metrics
	}

	// get version
	version, err := monitor.GetVersion(c.config)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.solanaVersion, err)
	} else {
		ch <- prometheus.MustNewConstMetric(c.solanaVersion, prometheus.GaugeValue, 1, version.Result.SolanaCore)
	}

	// get balance
	bal, err := monitor.GetBalance(c.config)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.accountBalance, err)
	} else {
		s := strconv.FormatInt(bal.Result.Value, 10)
		ch <- prometheus.MustNewConstMetric(c.accountBalance, prometheus.GaugeValue, 1, s)
	}

	// get slot leader
	leader, err := monitor.GetSlotLeader(c.config)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.slotLeader, err)
	} else {
		ch <- prometheus.MustNewConstMetric(c.slotLeader, prometheus.GaugeValue, 1, leader.Result)
	}

	// get current slot
	slot, err := monitor.GetCurrentSlot(c.config)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.currentSlot, err)
	} else {
		cs := strconv.FormatInt(slot.Result, 10)
		ch <- prometheus.MustNewConstMetric(c.currentSlot, prometheus.GaugeValue, 1, cs)
	}

	// get block time and calculate block time diff
	bt, err := monitor.GetBlockTime(slot.Result, c.config)
	if err != nil {
		log.Printf("Error while getting block time: %v", err)
	}

	// get previous block time
	pvt, err := monitor.GetBlockTime(slot.Result-1, c.config)
	if err != nil {
		log.Printf("Error while getting previous block time: %v", err)
	}

	// block tim difference
	sec, s := blockTimeDiff(bt.Result, pvt.Result)
	ch <- prometheus.MustNewConstMetric(c.blockTime, prometheus.GaugeValue, sec, s+"s")

	// IP address of gossip
	address := c.getClusterNodeInfo()
	ch <- prometheus.MustNewConstMetric(c.ipAddress, prometheus.GaugeValue, 1, address)

	// get tx count
	count, _ := monitor.GetTxCount(c.config)
	txcount := utils.NearestThousandFormat(float64(count.Result))

	ch <- prometheus.MustNewConstMetric(c.txCount, prometheus.GaugeValue, float64(count.Result), txcount)

}

func (c *solanaCollector) getClusterNodeInfo() string {
	result, err := monitor.GetClusterNodes(c.config)
	if err != nil {
		log.Printf("Error while getting cluster node information : %v", err)
	}
	var address string
	for _, value := range result.Result {
		if value.Pubkey == c.config.ValDetails.PubKey {
			// ch <- prometheus.MustNewConstMetric(c.ipAddress, prometheus.GaugeValue, 1, value.Gossip)
			address = value.Gossip
		}
	}
	return address
}

func (c *solanaCollector) getNetworkVoteAccountinfo() float64 {
	resn, _ := monitor.GetVoteAccounts(c.config, utils.Network)
	var outN float64
	for _, vote := range resn.Result.Current {
		if vote.NodePubkey == c.config.ValDetails.PubKey {
			outN = float64(vote.LastVote)

		}
	}
	return outN
}

func blockTimeDiff(bt int64, pvt int64) (float64, string) {
	t1 := time.Unix(bt, 0)
	t2 := time.Unix(pvt, 0)

	sub := t1.Sub(t2)
	diff := sub.Seconds()

	if diff < 0 {
		diff = -(diff)
	}
	s := fmt.Sprintf("%.2f", diff)

	sec, _ := strconv.ParseFloat(s, 64)

	return sec, s
}
