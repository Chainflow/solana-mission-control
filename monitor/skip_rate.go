package monitor

import (
	"encoding/json"
	"log"
	"os/exec"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/types"
)

func SkipRate(cfg *config.Config) (float64, error) {
	var skipedSlots float64

	cmd := exec.Command("solana", "validators", "--output", "json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while running ping command %v", err)
		return skipedSlots, err
	}

	var result types.SkipRate
	err = json.Unmarshal(out, &result)
	if err != nil {
		log.Printf("Error: %v", err)
		return skipedSlots, err
	}

	for _, val := range result.Validators {
		if val.IdentityPubkey == cfg.ValDetails.PubKey {
			skipedSlots = val.SkipRate
		}
	}

	log.Printf("VAL SKIPPED RATE : %f", skipedSlots)

	return skipedSlots, nil
}
