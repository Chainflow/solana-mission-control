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

	Commitment struct {
		Commitemnt string `json:"commitment"`
	}

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

	SolanaVersion struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  struct {
			SolanaCore string `json:"solana-core"`
		} `json:"result"`
		ID int `json:"id"`
	}

	LeaderShedule struct {
		Jsonrpc string             `json:"jsonrpc"`
		Result  map[string][]int64 `json:"result"`
		ID      int                `json:"id"`
	}

	ConfirmedBlocks struct {
		Jsonrpc string  `json:"jsonrpc"`
		Result  []int64 `json:"result"`
		ID      int     `json:"id"`
	}

	BlockTime struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  int64  `json:"result"`
		// ID int `json:"id"`
	}

	VoteAccounts struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  struct {
			Current []struct {
				Commission       int  `json:"commission"`
				EpochVoteAccount bool `json:"epochVoteAccount"`
				// EpochCredits     [][]int `json:"epochCredits"`
				NodePubkey     string `json:"nodePubkey"`
				LastVote       int    `json:"lastVote"`
				ActivatedStake int    `json:"activatedStake"`
				VotePubkey     string `json:"votePubkey"`
			} `json:"current"`
			Delinquent []struct {
				Commission       int  `json:"commission"`
				EpochVoteAccount bool `json:"epochVoteAccount"`
				// EpochCredits     []interface{} `json:"epochCredits"`
				NodePubkey     string `json:"nodePubkey"`
				LastVote       int    `json:"lastVote"`
				ActivatedStake int    `json:"activatedStake"`
				VotePubkey     string `json:"votePubkey"`
			} `json:"delinquent"`
		} `json:"result"`
		ID int `json:"id"`
	}

	NodeHealth struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  string `json:"result"`
		Error   struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    struct {
			} `json:"data"`
		} `json:"error"`
		// ID int `json:"id"`
	}

	Version struct {
		// Jsonrpc string `json:"jsonrpc"`
		Result struct {
			SolanaCore string `json:"solana-core"`
		} `json:"result"`
		// ID int `json:"id"`
	}

	Identity struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  struct {
			Identity string `json:"identity"`
		} `json:"result"`
		// ID int `json:"id"`
	}

	VoteAccount struct {
		ActivatedStake   int64   `json:"activatedStake"`
		Commission       int64   `json:"commission"`
		EpochCredits     [][]int `json:"epochCredits"`
		EpochVoteAccount bool    `json:"epochVoteAccount"`
		LastVote         int     `json:"lastVote"`
		NodePubkey       string  `json:"nodePubkey"`
		RootSlot         int     `json:"rootSlot"`
		VotePubkey       string  `json:"votePubkey"`
	}

	GetVoteAccountsResponse struct {
		Result struct {
			Current    []VoteAccount `json:"current"`
			Delinquent []VoteAccount `json:"delinquent"`
		} `json:"result"`
		Error rpcError `json:"error"`
	}

	rpcError struct {
		Message string `json:"message"`
		Code    int64  `json:"id"`
	}

	Stake struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  struct {
			Active   int64  `json:"active"`
			Inactive int64  `json:"inactive"`
			State    string `json:"state"`
		} `json:"result"`
		// ID int `json:"id"`
	}

	SlotLeader struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  string `json:"result"`
		// ID      int    `json:"id"`
	}

	CurrentSlot struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  int64  `json:"result"`
		// ID      int    `json:"id"`
	}

	DBRes struct {
		Status string `json:"status"`
		Data   struct {
			ResultType string `json:"resultType"`
			Result     []struct {
				Metric struct {
					Name             string `json:"__name__"`
					Instance         string `json:"instance"`
					Job              string `json:"job"`
					SolanaAccBalance string `json:"solana_acc_balance"`
					AlertCount       string `json:"alert_count"`
				} `json:"metric"`
				Value []interface{} `json:"value"`
			} `json:"result"`
		} `json:"data"`
	}

	TxCount struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  int64  `json:"result"`
		// ID      int    `json:"id"`
	}

	ClustrNode struct {
		// Jsonrpc string `json:"jsonrpc"`
		Result []struct {
			Gossip  string `json:"gossip"`
			Pubkey  string `json:"pubkey"`
			RPC     string `json:"rpc"`
			Tpu     string `json:"tpu"`
			Version string `json:"version"`
		} `json:"result"`
		// ID int `json:"id"`
	}
)
