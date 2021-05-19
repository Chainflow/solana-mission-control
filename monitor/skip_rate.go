package monitor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/types"
)

func SkipRate(cfg *config.Config) (float64, error) {
	// cmd := exec.Command("solana", "validators", "--output", "json")
	// // log.Fatalf("solana cmd..", cmd)
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Printf("Error while running ping command %v", err)
	// 	// return
	// }

	endPoint := fmt.Sprintf("https://www.validators.app/api/v1/validators/mainnet/%s.json", cfg.ValDetails.PubKey)
	ops := types.HTTPOptions{
		Endpoint: endPoint,
		// Endpoint: "https://www.validators.app/api/v1/validators/mainnet/7Gjec4iDbTxLvVYNsRbZrrHdtyLByzdDJ1C5BmcMMBks.json",
		Method: http.MethodGet,
	}
	var skipedSlots float64

	var result types.ValidatorsAPIResp
	resp, err := HitValditorsAPI(ops, cfg)
	if err != nil {
		log.Printf("Error: %v", err)
		return skipedSlots, err
	}

	err = json.Unmarshal(out, &result)
	if err != nil {
		log.Printf("Error: %v", err)
		return skipedSlots, err
	}
	skipedSlots = float64(result.SkippedSlots)

	log.Printf("skipped slots : %d", result.SkippedSlots)

	return skipedSlots, nil
}
