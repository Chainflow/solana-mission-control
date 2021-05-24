package types

import (
	client "github.com/influxdata/influxdb1-client/v2"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
)

type (
	// QueryParams map of strings
	QueryParams map[string]string

	// HTTPOptions is a structure that holds all http options parameters
	HTTPOptions struct {
		Endpoint    string
		QueryParams QueryParams
		Body        Payload
		Method      string
	}

	// Payload is a structure which holds all the curl payload parameters
	Payload struct {
		Jsonrpc string        `json:"jsonrpc"`
		Method  string        `json:"method"`
		Params  []interface{} `json:"params"`
		ID      int           `json:"id"`
	}

	// Commitement struct holds the state of Commitment
	Commitment struct {
		Commitemnt string `json:"commitment"`
	}
	// Encode struct to encode string
	Encode struct {
		Encoding string `json:"encoding"`
	}

	// Params struct
	Params struct {
		To     string `json:"to"`
		Data   string `json:"data"`
		Encode Encode `json:"encode"`
	}

	// Target is a structure which holds all the parameters of a target
	//this could be used to write endpoints for each functionality
	Target struct {
		ExecutionType string
		HTTPOptions   HTTPOptions
		Name          string
		Func          func(m HTTPOptions, cfg *config.Config, c client.Client)
		ScraperRate   string
	}

	// Targets list of all the targets
	Targets struct {
		List []Target
	}

	// PingResp is a structure which holds the options of a response
	PingResp struct {
		StatusCode int
		Body       []byte
	}

	// Balance struct which holds information of Account Balancce
	Balance struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  struct {
			Context struct {
				Slot int `json:"slot"`
			} `json:"context"`
			Value int64 `json:"value"`
		} `json:"result"`
		ID int `json:"id"`
	}

	// AccountInfo struct which holds Account Information
	AccountInfo struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  struct {
			Context struct {
				Slot int `json:"slot"`
			} `json:"context"`
			Value struct {
				Data struct {
					Nonce struct {
						Initialized struct {
							Authority     string `json:"authority"`
							Blockhash     string `json:"blockhash"`
							FeeCalculator struct {
								LamportsPerSignature int `json:"lamportsPerSignature"`
							} `json:"feeCalculator"`
						} `json:"initialized"`
					} `json:"nonce"`
				} `json:"data"`
				Executable bool   `json:"executable"`
				Lamports   int    `json:"lamports"`
				Owner      string `json:"owner"`
				RentEpoch  int    `json:"rentEpoch"`
			} `json:"value"`
		} `json:"result"`
		ID int `json:"id"`
	}

	// EpochInfo struct which holds information of current Epoch
	EpochInfo struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  struct {
			AbsoluteSlot int64 `json:"absoluteSlot"`
			BlockHeight  int64 `json:"blockHeight"`
			Epoch        int64 `json:"epoch"`
			SlotIndex    int64 `json:"slotIndex"`
			SlotsInEpoch int64 `json:"slotsInEpoch"`
		} `json:"result"`
		ID int `json:"id"`
	}

	// EpochShedule struct holds Epoch Shedule Information
	EpochShedule struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  struct {
			FirstNormalEpoch         int  `json:"firstNormalEpoch"`
			FirstNormalSlot          int  `json:"firstNormalSlot"`
			LeaderScheduleSlotOffset int  `json:"leaderScheduleSlotOffset"`
			SlotsPerEpoch            int  `json:"slotsPerEpoch"`
			Warmup                   bool `json:"warmup"`
		} `json:"result"`
		ID int `json:"id"`
	}

	// LeaderShedule struct holds information of leader schedule for an epoch
	LeaderShedule struct {
		Jsonrpc string             `json:"jsonrpc"`
		Result  map[string][]int64 `json:"result"`
		ID      int                `json:"id"`
	}

	// ConfirmedBlocks struct which holds information of confirmed blocks
	ConfirmedBlocks struct {
		Jsonrpc string  `json:"jsonrpc"`
		Result  []int64 `json:"result"`
		ID      int     `json:"id"`
	}

	// BlockTime struct which holds information of estimated production time of a confirmed block
	BlockTime struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  int64  `json:"result"`
	}

	// NodeHealth struct which holds information of health of the node
	NodeHealth struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  string `json:"result"`
		Error   struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    struct {
			} `json:"data"`
		} `json:"error"`
	}

	// Version struct which holds information of solana version
	Version struct {
		// Jsonrpc string `json:"jsonrpc"`
		Result struct {
			SolanaCore string `json:"solana-core"`
		} `json:"result"`
	}

	// Identity struct holds the pubkey for the current node
	Identity struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  struct {
			Identity string `json:"identity"`
		} `json:"result"`
	}

	// VoteAccount struct holds information of vote account
	VoteAccount struct {
		ActivatedStake   int64     `json:"activatedStake"`
		Commission       int64     `json:"commission"`
		EpochCredits     [][]int64 `json:"epochCredits"`
		EpochVoteAccount bool      `json:"epochVoteAccount"`
		LastVote         int       `json:"lastVote"`
		NodePubkey       string    `json:"nodePubkey"`
		RootSlot         int       `json:"rootSlot"`
		VotePubkey       string    `json:"votePubkey"`
	}

	// GetVoteAccountsResponse struct holds information of current and deliquent vote accounts
	GetVoteAccountsResponse struct {
		Result struct {
			Current    []VoteAccount `json:"current"`
			Delinquent []VoteAccount `json:"delinquent"`
		} `json:"result"`
		Error rpcError `json:"error"`
	}

	// rpcError struct which holds Error message of RPC
	rpcError struct {
		Message string `json:"message"`
		Code    int64  `json:"id"`
	}
	// Stake struct which holds information of stake account
	Stake struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  struct {
			Active   int64  `json:"active"`
			Inactive int64  `json:"inactive"`
			State    string `json:"state"`
		} `json:"result"`
	}

	// SlotLeader holds the  information of current slot leader
	SlotLeader struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  string `json:"result"`
	}

	// CurrentSlot holds the information of Current slot
	CurrentSlot struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  int64  `json:"result"`
	}

	// DBRes struct holds the Account balance and alertcount which stored in Database
	DBRes struct {
		Status string `json:"status"`
		Data   struct {
			ResultType string `json:"resultType"`
			Result     []struct {
				Metric struct {
					Name                  string `json:"__name__"`
					Instance              string `json:"instance"`
					Job                   string `json:"job"`
					SolanaAccBalance      string `json:"solana_acc_balance"`
					AlertCount            string `json:"alert_count"`
					SolanaValStatus       string `json:"solana_val_status"`
					SolanaPreviousCredits string `json:"solana_previous_credits"`
					SolanaCurrentCredits  string `json:"solana_current_credits"`
				} `json:"metric"`
				Value []interface{} `json:"value"`
			} `json:"result"`
		} `json:"data"`
	}

	// TxCount struct which holds information of Transaction count
	TxCount struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  int64  `json:"result"`
	}

	// ClusterNode struct which holds information about all the nodes participating in the cluster
	ClustrNode struct {
		// Jsonrpc string `json:"jsonrpc"`
		Result []struct {
			Gossip  string `json:"gossip"`
			Pubkey  string `json:"pubkey"`
			RPC     string `json:"rpc"`
			Tpu     string `json:"tpu"`
			Version string `json:"version"`
		} `json:"result"`
	}

	// ConfirmedBlock struct which holds blocktime of confirmedBlock at current slot height
	ConfirmedBlock struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  struct {
			BlockTime int64 `json:"blockTime"`
		} `json:"result"`
	}

	ValidatorsAPIResp struct {
		Network                      string      `json:"network"`
		Account                      string      `json:"account"`
		Name                         string      `json:"name"`
		KeybaseID                    string      `json:"keybase_id"`
		WwwURL                       string      `json:"www_url"`
		Details                      string      `json:"details"`
		CreatedAt                    string      `json:"created_at"`
		UpdatedAt                    string      `json:"updated_at"`
		TotalScore                   int         `json:"total_score"`
		RootDistanceScore            int         `json:"root_distance_score"`
		VoteDistanceScore            int         `json:"vote_distance_score"`
		SkippedSlotScore             int         `json:"skipped_slot_score"`
		SoftwareVersion              string      `json:"software_version"`
		SoftwareVersionScore         int         `json:"software_version_score"`
		StakeConcentrationScore      int         `json:"stake_concentration_score"`
		DataCenterConcentrationScore int         `json:"data_center_concentration_score"`
		PublishedInformationScore    int         `json:"published_information_score"`
		SecurityReportScore          int         `json:"security_report_score"`
		ActiveStake                  int64       `json:"active_stake"`
		Commission                   int         `json:"commission"`
		Delinquent                   bool        `json:"delinquent"`
		DataCenterKey                string      `json:"data_center_key"`
		DataCenterHost               interface{} `json:"data_center_host"`
		VoteAccount                  string      `json:"vote_account"`
		SkippedSlots                 int         `json:"skipped_slots"`
		SkippedSlotPercent           string      `json:"skipped_slot_percent"`
		PingTime                     interface{} `json:"ping_time"`
		URL                          string      `json:"url"`
	}

	SkipRate struct {
		TotalActiveStake     int64 `json:"totalActiveStake"`
		TotalCurrentStake    int64 `json:"totalCurrentStake"`
		TotalDelinquentStake int64 `json:"totalDelinquentStake"`
		Validators           []struct {
			IdentityPubkey    string  `json:"identityPubkey"`
			VoteAccountPubkey string  `json:"voteAccountPubkey"`
			Commission        int     `json:"commission"`
			LastVote          int     `json:"lastVote"`
			RootSlot          int     `json:"rootSlot"`
			Credits           int     `json:"credits"`
			EpochCredits      int     `json:"epochCredits"`
			ActivatedStake    int64   `json:"activatedStake"`
			Version           string  `json:"version"`
			Delinquent        bool    `json:"delinquent"`
			SkipRate          float64 `json:"skipRate"`
		} `json:"validators"`
		StakeByVersion interface{} `json:"stakeByVersion"`
	}

	BlockProduction struct {
		Epoch               int `json:"epoch"`
		StartSlot           int `json:"start_slot"`
		EndSlot             int `json:"end_slot"`
		TotalSlots          int `json:"total_slots"`
		TotalBlocksProduced int `json:"total_blocks_produced"`
		TotalSlotsSkipped   int `json:"total_slots_skipped"`
		Leaders             []struct {
			IdentityPubkey string `json:"identityPubkey"`
			LeaderSlots    int    `json:"leaderSlots"`
			BlocksProduced int    `json:"blocksProduced"`
			SkippedSlots   int    `json:"skippedSlots"`
		} `json:"leaders"`
		// IndividualSlotStatus []struct {
		// 	Slot    int    `json:"slot"`
		// 	Leader  string `json:"leader"`
		// 	Skipped bool   `json:"skipped"`
		// } `json:"individual_slot_status"`
	}
)
