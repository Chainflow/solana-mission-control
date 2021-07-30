package monitor_test

import (
	"testing"

	"github.com/Chainflow/solana-mission-control/config"
	"github.com/Chainflow/solana-mission-control/monitor"
	"github.com/Chainflow/solana-mission-control/utils"
)

func TestConfirmedBlocks(t *testing.T) {
	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config :", err)
	}
	res, err := monitor.GetConfirmedBlocks(70539000, 70539097, cfg)
	if err != nil {
		t.Error("Error while getting confirmed blocks at given range")
	}
	if &res == nil {
		t.Error("Expected non empty result, but got empty result: ", res)
	}
	if &res != nil {
		t.Log("Got confirmed Blocks between given range : ", res)
	}
}

func TestConfirmedValidatorBlock(t *testing.T) {

	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config :", err)
	}
	res, err := monitor.GetConfirmedBlock(cfg, 70539097, utils.Validator)
	if err != nil {
		t.Error("Error while getting confirmed block Information at given height")
	}
	if &res.Result == nil {
		t.Error("Expected non empty result, but got empty result: ", res)
	}
	if &res.Result != nil {
		t.Log("Got confirmed Block information : ", res.Result)
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
		t.Error("Expected non empty result, but got empty result: ", res)
	}
	if res.Result.BlockTime != 0 {
		t.Log("Got network confirmed Block information : ", res)
	}
}
