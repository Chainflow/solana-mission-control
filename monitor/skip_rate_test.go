package monitor_test

import (
	"testing"

	"github.com/Chainflow/solana-mission-control/config"
	"github.com/Chainflow/solana-mission-control/monitor"
)

func TestSkipRate(t *testing.T) {

	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config :", err)
	}
	vs, ns, err := monitor.SkipRate(cfg)
	if err != nil {
		t.Error("Error while fetching skip rate")
	}
	if &vs == nil {
		t.Error("Expected non empty result, but got empty result: ", vs)
	}
	if &ns == nil {
		t.Error("Expected non empty result, but got empty result: ", ns)
	}
	if &vs != nil {
		t.Log("Got validatator skip rate", vs)
	}
	if &ns != nil {
		t.Log("Got network skip rate", ns)
	}

}
