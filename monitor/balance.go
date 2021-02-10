package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	// "k8s.io/klog/v2"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/types"
)

func GetBalance(cfg *config.Config) (types.Balance, error) {
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.RPCEndpoint,
		Method:   http.MethodPost,
		Body: types.Payload{Jsonrpc: "2.0", Method: "getBalance", ID: 1, Params: []interface{}{
			cfg.ValDetails.PubKey, // should be base58 encoded to query data
			// "83astBRguLMdt2h5U1Tpdq5tjFoJ6noeGwaY3mDLVcri",
			// types.Encode{
			// 	Encoding: "jsonParsed",
			// },
		}},
	}

	var result types.Balance
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

func GetIdentity(cfg *config.Config) (types.Identity, error) {
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.RPCEndpoint,
		Method:   http.MethodPost,
		Body:     types.Payload{Jsonrpc: "2.0", Method: "getIdentity", ID: 1},
	}

	var result types.Identity
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
