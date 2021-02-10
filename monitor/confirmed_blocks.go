package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/types"
)

func GetConfirmedBlocks(rangeStart int64, rangeEnd int64, cfg *config.Config) ([]int64, error) {
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.RPCEndpoint,
		Method:   http.MethodPost,
		Body:     types.Payload{Jsonrpc: "2.0", Method: "getConfirmedBlocks", ID: 1, Params: []interface{}{rangeStart, rangeEnd}},
	}

	// log.Fatalf("ops...", ops)

	// cfm, err := scraper.GetConfirmedBlocks(ops)
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
