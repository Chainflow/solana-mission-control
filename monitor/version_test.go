package monitor_test

import (
	"testing"

	"github.com/Chainflow/solana-mission-control/config"
	"github.com/Chainflow/solana-mission-control/monitor"
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

	if &res.Result == nil {
		t.Error("Expected non empty result, but got empty version : ", res.Result.SolanaCore)
	}

	if &res.Result != nil {
		t.Log("Got Version : ", res.Result.SolanaCore)
	}
}
