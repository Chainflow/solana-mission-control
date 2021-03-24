package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/types"
	"github.com/PrathyushaLakkireddy/solana-prometheus/utils"
)

// GetConfirmedBlocks returns a list of confirmed blocks between two slots of given range.
func GetConfirmedBlocks(rangeStart int64, rangeEnd int64, cfg *config.Config) ([]int64, error) {
	log.Println("Getting Confirmed Blocks...")
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.RPCEndpoint,
		Method:   http.MethodPost,
		Body:     types.Payload{Jsonrpc: "2.0", Method: "getConfirmedBlocks", ID: 1, Params: []interface{}{rangeStart, rangeEnd}},
	}

	var cfm types.ConfirmedBlocks
	resp, err := HitHTTPTarget(ops)
	if err != nil {
		log.Printf("Error while getting leader shedules: %v", err)
		return nil, err
	}

	err = json.Unmarshal(resp.Body, &cfm)
	if err != nil {
		log.Printf("Error while unmarshelling leader shedules: %v", err)
		return nil, err
	}
	return cfm.Result, nil
}

// GetConfirmedBlock takes current slot height and returns identity and transaction information about a
// confirmed block in the ledger
func GetConfirmedBlock(cfg *config.Config, height int64, node string) (types.ConfirmedBlock, error) {
	log.Println("Getting Confirmed Block...")
	var result types.ConfirmedBlock
	ops := types.HTTPOptions{
		Method: http.MethodPost,
		Body:   types.Payload{Jsonrpc: "2.0", Method: "getConfirmedBlock", ID: 1, Params: []interface{}{height}},
	}

	if node == utils.Network {
		ops.Endpoint = cfg.Endpoints.NetworkRPC
	} else if node == utils.Validator {
		ops.Endpoint = cfg.Endpoints.RPCEndpoint
	} else {
		ops.Endpoint = cfg.Endpoints.RPCEndpoint
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
