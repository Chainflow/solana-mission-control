package monitor

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/Chainflow/solana-mission-control/alerter"
	"github.com/Chainflow/solana-mission-control/config"
	"github.com/Chainflow/solana-mission-control/types"
	"github.com/Chainflow/solana-mission-control/utils"
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

	voteAccounts, err := GetVoteAccounts(cfg, utils.Network)
	if err != nil {
		log.Printf("Error while getting vote accounts : %v", err)
	}

	if &voteAccounts.Result != nil {
		currentVal := len(voteAccounts.Result.Current)

		netSkipped = totalSkipped / float64(currentVal)
	}

	log.Printf("VAL skip rate : %f, Network skip rate : %f", valSkipped, netSkipped)

	if valSkipped > netSkipped {
		if strings.EqualFold(cfg.AlerterPreferences.SkipRateAlerts, "yes") {
			err = alerter.SendTelegramAlert(fmt.Sprintf("SKIP RATE ALERT ::  Your validator SKIP RATE : %f has exceeded network SKIP RATE : %f", valSkipped, netSkipped), cfg)
			if err != nil {
				log.Printf("Error while sending skip rate alert to telegram: %v", err)
			}
			err = alerter.SendEmailAlert(fmt.Sprintf("Your validator SKIP RATE : %f has exceeded network SKIP RATE : %f", valSkipped, netSkipped), cfg)
			if err != nil {
				log.Printf("Error while sending skip rate alert to email: %v", err)
			}
		}
	}

	return valSkipped, netSkipped, nil
}
