package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/types"
)

// GetBlockTime returns the estimated production time of a confirmed block
func GetBlockTime(slot int64, cfg *config.Config) (types.BlockTime, error) {
	log.Println("Getting block time...")
	var result types.BlockTime
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.RPCEndpoint,
		Method:   http.MethodPost,
		Body:     types.Payload{Jsonrpc: "2.0", Method: "getBlockTime", ID: 1, Params: []interface{}{slot}},
	}

	resp, err := HitHTTPTarget(ops)
	if err != nil {
		log.Printf("Error while getting block time: %v", err)
		return result, err
	}

	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		log.Printf("Error while unmarshelling block time res: %v", err)
		return result, err
	}
	return result, nil
}
