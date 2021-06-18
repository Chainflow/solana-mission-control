package monitor_test

import (
	"testing"

	"github.com/Chainflow/solana-mission-control/config"
	"github.com/Chainflow/solana-mission-control/monitor"
	"github.com/Chainflow/solana-mission-control/utils"
)

func TestGetValidatorEpochinfo(t *testing.T) {
	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config :", err)
	}
	res, err := monitor.GetEpochInfo(cfg, utils.Validator)
	if err != nil {
		t.Error("Error while fetching Epoch Information")
	}
	if &res.Result == nil {
		t.Error("Expected non empty result, but got empty result: ", res.Result)
	}
	if &res.Result != nil {
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
		t.Error("Error while fetching epoch information")
	}
	if &res.Result == nil {
		t.Error("Expected non empty result, but got empty result: ", res.Result)
	}
	if &res.Result.Epoch != nil {
		t.Log("Got Network Epoch information: ", res.Result)
	}
}
