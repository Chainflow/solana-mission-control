package monitor_test

import (
	"testing"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/monitor"
	"github.com/PrathyushaLakkireddy/solana-prometheus/utils"
)

func TestGetValidatorVoteAccounts(t *testing.T) {
	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config :", err)
	}
	res, err := monitor.GetVoteAccounts(cfg, utils.Validator)
	if err != nil {
		t.Error("Error while Validator Vote Account information")
	}
	if &res.Result.Current == nil || &res.Result.Delinquent == nil {
		t.Error("Expected non empty result, but got empty result: ", res.Result)
	}
	if &res.Result.Current != nil {
		t.Log("Got Validator current vote account information", res.Result)
	} else if &res.Result.Delinquent != nil {
		t.Log("Got Validator Deliquent vote account information", res.Result)
	}
}

func TestGetNetworkVoteAccounts(t *testing.T) {
	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config :", err)
	}
	res, err := monitor.GetVoteAccounts(cfg, utils.Network)
	if err != nil {
		t.Error("Error while fetching Network Vote Account Information")
	}
	if res.Result.Current == nil {
		t.Error("Expected non empty result, but got empty result: ", res.Result)
	} else if res.Result.Delinquent == nil {
		t.Error("Expected non empty result, but got empty result: ", res.Result)
	}
	if res.Result.Current != nil {
		t.Log("Got Network current vote account information", res.Result)
	} else if res.Result.Delinquent != nil {
		t.Log("Got Network Deliquent vote account information", res.Result)
	}
}
