package monitor_test

import (
	"testing"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/monitor"
)

func TestGetBalance(t *testing.T) {
	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config :", err)
	}
	res, err := monitor.GetBalance(cfg)
	if err != nil {
		t.Error("Error while fetching Account balance")
	}
	if res.Result.Value == 0 {
		t.Error("Expected non empty result, but got empty result: ", res.Result)
	}
	if res.Result.Value != 0 {
		t.Log("Got Account Balance: ", res.Result)
	}
}
