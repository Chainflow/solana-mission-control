package monitor_test

import (
	"testing"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/monitor"
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
	if res.Result == 0 {
		t.Error("Expexted non empty result, but got empty result: ", res.Result)
	}
	if res.Result != 0 {
		t.Log("Got block time: ", res.Result)
	}
}
