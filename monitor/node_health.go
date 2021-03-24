package monitor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PrathyushaLakkireddy/solana-prometheus/alerter"
	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/types"
)

// GetNodeHealth returns the current health of the node.
func GetNodeHealth(cfg *config.Config) (types.NodeHealth, error) {
	log.Println("Getting Node Health...")
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.RPCEndpoint,
		Method:   http.MethodPost,
		Body:     types.Payload{Jsonrpc: "2.0", Method: "getHealth", ID: 1},
	}

	var result types.NodeHealth
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

	if result.Result != "" {
		if strings.EqualFold(result.Result, "ok") {
			log.Printf("Node health : %s", result.Result)
		} else {
			err = alerter.SendTelegramAlert(fmt.Sprintf("Your node is not running"), cfg)
			if err != nil {
				log.Printf("Error while sending node health alert: %v", err)
			}
		}
	}

	return result, nil
}
