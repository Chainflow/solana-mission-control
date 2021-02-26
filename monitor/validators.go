package monitor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/types"
)

func GetVoteAccounts(cfg *config.Config) (types.GetVoteAccountsResponse, error) {
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.RPCEndpoint,
		Method:   http.MethodPost,
		Body: types.Payload{Jsonrpc: "2.0", Method: "getVoteAccounts", ID: 1, Params: []interface{}{
			types.Commitment{
				Commitemnt: "recent",
			},
		}},
	}

	var result types.GetVoteAccountsResponse

	resp, err := HitHTTPTarget(ops)
	if err != nil {
		log.Printf("Error while getting leader shedules: %v", err)
		return result, err
	}

	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		log.Printf("Error while unmarshelling leader shedules: %v", err)
		return result, err
	}

	if result.Error.Code != 0 {
		return result, fmt.Errorf("RPC error: %d %v", result.Error.Code, result.Error.Message)
	}

	return result, nil
}
