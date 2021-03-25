package monitor_test

import (
	"testing"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/monitor"
)

func TestGetNodeHealth(t *testing.T) {
	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config :", err)
	}
	res, err := monitor.GetNodeHealth(cfg)
	if err != nil {
		t.Error("Error while fetching Node Health")
	}
	if res == 0 {
		t.Error("Expexted OK, but got empty result: ", res)
	}
	if res != 0 {
		t.Log("Got Node Health", res)
	}
}
