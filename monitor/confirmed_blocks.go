package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/types"
	"github.com/PrathyushaLakkireddy/solana-prometheus/utils"
)

func GetConfirmedBlocks(rangeStart int64, rangeEnd int64, cfg *config.Config) ([]int64, error) {
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

func GetConfirmedBlock(cfg *config.Config, height int64, node string) (types.ConfirmedBlock, error) {
	var result types.ConfirmedBlock
	ops := types.HTTPOptions{
		// Endpoint: cfg.Endpoints.RPCEndpoint,
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
