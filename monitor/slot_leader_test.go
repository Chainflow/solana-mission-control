package monitor_test

import (
	"testing"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/monitor"
)

func TestGetSlotLeader(t *testing.T) {
	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config :", err)
	}
	res, err := monitor.GetSlotLeader(cfg)
	if err != nil {
		t.Error("Error while fetching Slot Leader")
	}
	if res.Result == "" {
		t.Error("Expexted non empty result, but got empty result: ", res.Result)
	}
	if res.Result != "" {
		t.Log("Got Slot Leader", res.Result)
	}
}
