package monitor

import (
	"encoding/json"
	"log"
	"os/exec"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/types"
)

func SkipRate(cfg *config.Config) (float64, float64, error) {
	var valSkipped, netSkipped, totalSkipped float64

	cmd := exec.Command("solana", "validators", "--output", "json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while running solana validators cli command %v", err)
		return valSkipped, netSkipped, err
	}

	var result types.SkipRate
	err = json.Unmarshal(out, &result)
	if err != nil {
		log.Printf("Error: %v", err)
		return valSkipped, netSkipped, err
	}

	for _, val := range result.Validators {
		if val.IdentityPubkey == cfg.ValDetails.PubKey {
			valSkipped = val.SkipRate
		}
		totalSkipped = totalSkipped + val.SkipRate
	}
	netSkipped = totalSkipped / float64(len(result.Validators))

	log.Printf("VAL skip rate : %f, Network skip rate : %f", valSkipped, netSkipped)

	return valSkipped, netSkipped, nil
}
