package monitor_test

import (
	"testing"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/monitor"
	"github.com/PrathyushaLakkireddy/solana-prometheus/utils"
)

func TestConfirmedValidatorBlock(t *testing.T) {

	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config :", err)
	}
	res, err := monitor.GetConfirmedBlock(cfg, 70539097, utils.Validator)
	if err != nil {
		t.Error("Error while getting confirmed block Information at given height")
	}
	if res.Result.BlockTime == 0 {
		t.Error("Expexted non empty result, but got empty result: ", res)
	}
	if res.Result.BlockTime != 0 {
		t.Log("Got confirmed Block information : ", res)
	}
}

func TestConfirmedNetworkBlock(t *testing.T) {

	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config :", err)
	}
	res, err := monitor.GetConfirmedBlock(cfg, 70545876, utils.Network)
	if err != nil {
		t.Error("Error while getting network confirmed block Information at given heights")
	}
	if res.Result.BlockTime == 0 {
		t.Error("Expexted non empty result, but got empty result: ", res)
	}
	if res.Result.BlockTime != 0 {
		t.Log("Got network confirmed Block information : ", res)
	}
}
