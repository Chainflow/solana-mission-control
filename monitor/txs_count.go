package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Chainflow/solana-mission-control/config"
	"github.com/Chainflow/solana-mission-control/types"
)

// GetTxCount returns the current Transaction count from the ledger
func GetTxCount(cfg *config.Config) (types.TxCount, error) {
	var result types.TxCount
	var err error
	if cfg == nil {
		return result, err
	}
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.RPCEndpoint,
		Method:   http.MethodPost,
		Body:     types.Payload{Jsonrpc: "2.0", Method: "getTransactionCount", ID: 1},
	}

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
