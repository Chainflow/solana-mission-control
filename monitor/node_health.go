package monitor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Chainflow/solana-mission-control/alerter"
	"github.com/Chainflow/solana-mission-control/config"
	"github.com/Chainflow/solana-mission-control/types"
)

// GetNodeHealth returns the current health of the node.
func GetNodeHealth(cfg *config.Config) (float64, error) {
	log.Println("Getting Node Health...")
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.RPCEndpoint,
		Method:   http.MethodPost,
		Body:     types.Payload{Jsonrpc: "2.0", Method: "getHealth", ID: 1},
	}
	var h float64

	var result types.NodeHealth
	resp, err := HitHTTPTarget(ops)
	if err != nil {
		log.Printf("Error: %v", err)
		return h, err
	}

	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		log.Printf("Error: %v", err)
		return h, err
	}

	// send alert if node is down
	if result.Result != "" {
		if strings.EqualFold(result.Result, "ok") {
			log.Printf("Node health : %s", result.Result)
			h = 1

			return h, nil
		} else {
			if strings.EqualFold(cfg.AlerterPreferences.NodeHealthAlert, "yes") {
				err = alerter.SendTelegramAlert(fmt.Sprintf("Your node is not running"), cfg)
				if err != nil {
					log.Printf("Error while sending node health alert to telegram: %v", err)
				}
				err = alerter.SendEmailAlert(fmt.Sprintf("Your node is not running"), cfg)
				if err != nil {
					log.Printf("Error while sending node health alert to email: %v", err)
				}
				h = 0
			}
		}
	}

	return h, nil
}
