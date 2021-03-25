package monitor_test

import (
	"testing"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/monitor"
	"github.com/PrathyushaLakkireddy/solana-prometheus/utils"
)

func TestValidatorCurrentSlot(t *testing.T) {
	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config :", err)
	}

	res, err := monitor.GetCurrentSlot(cfg, utils.Validator)
	if err != nil {
		t.Error("Error while fetching validator current slot : ", err)
	}

	if res.Result == 0 {
		t.Error("Expexted non empty result, but got empty version : ", res.Result)
	}

	if res.Result != 0 {
		t.Log("Got validator current slot : ", res.Result)
	}
}

func TestNetworkCurrentSlot(t *testing.T) {
	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config :", err)
	}

	res, err := monitor.GetCurrentSlot(cfg, utils.Network)
	if err != nil {
		t.Error("Error while fetching network current slot : ", err)
	}

	if res.Result == 0 {
		t.Error("Expexted non empty result, but got empty result: ", res.Result)
	}

	if res.Result != 0 {
		t.Log("Got network current slot : ", res.Result)
	}
}