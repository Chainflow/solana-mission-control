package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/types"
)

// GetStake returns epoch activation information for a stake account
func GetStake(cfg *config.Config) (types.Stake, error) {
	log.Println("Getting Stake...")
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.RPCEndpoint,
		Method:   http.MethodPost,
		Body: types.Payload{Jsonrpc: "2.0", Method: "getBalance", ID: 1, Params: []interface{}{
			cfg.ValDetails.PubKey, // should be base58 encoded to query data
		}},
	}

	var result types.Stake
	resp, err := HitHTTPTarget(ops)
	if err != nil {
		log.Printf("Error: %v", err)
		return result, err
	}

	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		log.Printf("Error: %v", err)
		return result, err
	}

	return result, nil
}
