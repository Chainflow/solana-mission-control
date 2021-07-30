package monitor_test

import (
	"testing"

	"github.com/Chainflow/solana-mission-control/config"
	"github.com/Chainflow/solana-mission-control/monitor"
)

func TestBlockTime(t *testing.T) {
	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config :", err)
	}
	res, err := monitor.GetBlockTime(70539000, cfg)
	if err != nil {
		t.Error("Error while fetching block time")
	}
	if &res == nil {
		t.Error("Expected non empty result, but got empty result: ", res.Result)
	}
	if res.Result != 0 {
		t.Log("Got block time: ", res.Result)
	}
}
