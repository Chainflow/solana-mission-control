package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	// "k8s.io/klog/v2"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/types"
	"github.com/PrathyushaLakkireddy/solana-prometheus/utils"
)

func GetEpochInfo(cfg *config.Config, node string) (types.EpochInfo, error) {
	ops := types.HTTPOptions{
		// Endpoint: cfg.Endpoints.RPCEndpoint,
		Method: http.MethodPost,
		Body:   types.Payload{Jsonrpc: "2.0", Method: "getEpochInfo", ID: 1},
	}

	if node == utils.Network {
		ops.Endpoint = cfg.Endpoints.NetworkRPC
	} else if node == utils.Validator {
		ops.Endpoint = cfg.Endpoints.RPCEndpoint
	} else {
		ops.Endpoint = cfg.Endpoints.RPCEndpoint
	}

	var result types.EpochInfo
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
