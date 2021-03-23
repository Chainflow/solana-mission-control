package monitor_test

import (
	"testing"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/monitor"
)

func TestVersion(t *testing.T) {
	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config :", err)
	}

	res, err := monitor.GetVersion(cfg)
	if err != nil {
		t.Error("Error while fetching version : ", err)
	}

	if res.Result.SolanaCore == "" {
		t.Error("Expexted non empty result, but got empty version : ", res.Result.SolanaCore)
	}

	if res.Result.SolanaCore != "" {
		t.Log("Got Version : ", res.Result.SolanaCore)
	}
}
