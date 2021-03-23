package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/types"
)

// GetTxCount returns the current Transaction count from the ledger
func GetTxCount(cfg *config.Config) (types.TxCount, error) {
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.RPCEndpoint,
		Method:   http.MethodPost,
		Body:     types.Payload{Jsonrpc: "2.0", Method: "getTransactionCount", ID: 1},
	}

	var result types.TxCount
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

	// err = SendBalanceChangeAlert(result.Result.Value, cfg)
	// if err != nil {
	// 	log.Printf("Error while sending balance change alert : %v", err)
	// }

	return result, nil
}
