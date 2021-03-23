package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/types"
)

// GetClusterNodes returns information about all the nodes participating in the cluster
func GetClusterNodes(cfg *config.Config) (types.ClustrNode, error) {
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.RPCEndpoint,
		Method:   http.MethodPost,
		Body:     types.Payload{Jsonrpc: "2.0", Method: "getClusterNodes", ID: 1},
	}

	// if node == utils.Network {
	// 	ops.Endpoint = cfg.Endpoints.NetworkRPC
	// } else if node == utils.Validator {
	// 	ops.Endpoint = cfg.Endpoints.RPCEndpoint
	// } else {
	// 	ops.Endpoint = cfg.Endpoints.RPCEndpoint
	// }

	var result types.ClustrNode
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
