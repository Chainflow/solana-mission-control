package monitor_test

import (
	"testing"

	"github.com/Chainflow/solana-mission-control/config"
	"github.com/Chainflow/solana-mission-control/monitor"
)

func TestBlockProduction(t *testing.T) {
	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config :", err)
	}
	res, err := monitor.BlockProduction(cfg)
	if err != nil {
		t.Error("Error while fetching block production time")
	}
	if &res == nil {
		t.Error("Expected non empty result, but got empty result: ", res)
	}
	if &res != nil {
		t.Log("Got Block Production: ", res)
	}
}
