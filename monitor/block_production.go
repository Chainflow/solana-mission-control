package monitor

import (
	"encoding/json"
	"log"
	"os/exec"

	"github.com/Chainflow/solana-mission-control/config"
	"github.com/Chainflow/solana-mission-control/types"
)

type RecentBlock struct {
	TotalSlots          int `json:"total_slots"`
	TotalBlocksProduced int `json:"total_blocks_produced"`
	TotalSlotsSkipped   int `json:"total_slots_skipped"`
	LeaderSlots         int `json:"leaderSlots"`
	BlocksProduced      int `json:"blocksProduced"`
	SkippedSlots        int `json:"skippedSlots"`
}

func BlockProduction(cfg *config.Config) (RecentBlock, error) {
	var leaderSlots, blocksProduced, skippedSlots int
	var bp RecentBlock

	if solanaBinaryPath == "" {
		solanaBinaryPath = "solana"
	}

	log.Printf("Solana binary path : %s", solanaBinaryPath)

	cmd := exec.Command(solanaBinaryPath, "block-production", "--output", "json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while running solana validators cli command %v", err)
		return bp, err
	}

	var result types.BlockProduction
	err = json.Unmarshal(out, &result)
	if err != nil {
		log.Printf("Error: %v", err)
		return bp, err
	}

	for _, v := range result.Leaders {
		if v.IdentityPubkey == cfg.ValDetails.PubKey {
			leaderSlots = v.LeaderSlots
			blocksProduced = v.BlocksProduced
			skippedSlots = v.SkippedSlots
		}
	}

	bp = RecentBlock{
		TotalSlots:          result.TotalSlots,
		TotalBlocksProduced: result.TotalBlocksProduced,
		TotalSlotsSkipped:   result.TotalSlotsSkipped,
		LeaderSlots:         leaderSlots,
		BlocksProduced:      blocksProduced,
		SkippedSlots:        skippedSlots,
	}

	log.Printf("Block Production : %v", bp)

	return bp, nil
}
