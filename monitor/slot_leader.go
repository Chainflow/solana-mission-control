package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Chainflow/solana-mission-control/config"
	"github.com/Chainflow/solana-mission-control/types"
)

// GetSlotLeader returns the current slot leader
func GetSlotLeader(cfg *config.Config) (types.SlotLeader, error) {
	log.Println("Getting Slot Leader...")
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.RPCEndpoint,
		Method:   http.MethodPost,
		Body:     types.Payload{Jsonrpc: "2.0", Method: "getSlotLeader", ID: 1},
	}

	var result types.SlotLeader
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
