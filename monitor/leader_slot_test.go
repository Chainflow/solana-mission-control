package monitor_test

import (
	"testing"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/monitor"
)

func TestGetLeaderSlots(t *testing.T) {
	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config :", err)
	}
	res, err := monitor.GetLeaderSlots(163, cfg)
	if err != nil {
		t.Error("Error while fetching Leader Slot information")
	}
	if res == nil {
		t.Error("Expexted non empty result, but got empty result: ", res)
	}
	if res != nil {
		t.Log("Got Leader Slot Information: ", res)
	}
}
