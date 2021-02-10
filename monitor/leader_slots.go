package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/types"
)

func GetLeaderSlots(epochSlot int64, cfg *config.Config) (map[int64]string, error) {
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.RPCEndpoint,
		Method:   http.MethodPost,
		Body:     types.Payload{Jsonrpc: "2.0", Method: "getLeaderSchedule", ID: 1, Params: []interface{}{epochSlot}},
	}

	// log.Fatalf("ops...", ops, epochSlot)
	var sch types.LeaderShedule

	resp, err := HitHTTPTarget(ops)
	if err != nil {
		log.Printf("Error while getting leader shedules: %v", err)
		return nil, err
	}

	err = json.Unmarshal(resp.Body, &sch)
	if err != nil {
		log.Printf("Error while unmarshelling leader shedules: %v", err)
		return nil, err
	}

	slots := make(map[int64]string)

	for pk, sch := range sch.Result {
		if pk == cfg.ValDetails.PubKey {
			for _, i := range sch {
				slots[int64(i)] = pk
				// log.Printf("i,pk, sch", i, pk)
			}
		}
	}

	return slots, nil
}
