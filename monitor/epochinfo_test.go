package monitor_test

import (
	"testing"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/monitor"
	"github.com/PrathyushaLakkireddy/solana-prometheus/utils"
)

func TestGetValidatorEpochinfo(t *testing.T) {
	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config :", err)
	}
	res, err := monitor.GetEpochInfo(cfg, utils.Validator)
	if err != nil {
		t.Error("Error while fetching block time")
	}
	if res.Result.Epoch == 0 {
		t.Error("Expexted non empty result, but got empty result: ", res.Result)
	}
	if res.Result.Epoch != 0 {
		t.Log("Got Epoch information: ", res.Result)
	}
}

func TestGetNetworkEpochinfo(t *testing.T) {
	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config :", err)
	}
	res, err := monitor.GetEpochInfo(cfg, utils.Network)
	if err != nil {
		t.Error("Error while fetching block time")
	}
	if res.Result.Epoch == 0 {
		t.Error("Expexted non empty result, but got empty result: ", res.Result)
	}
	if res.Result.Epoch != 0 {
		t.Log("Got Network Epoch information: ", res.Result)
	}
}
