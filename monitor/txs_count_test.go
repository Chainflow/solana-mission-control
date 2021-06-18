package monitor

import (
	"testing"

	"github.com/Chainflow/solana-mission-control/config"
)

func TestTxCount(t *testing.T) {
	testCases := []struct {
		name    string
		before  func() *config.Config
		count   int64
		isError bool
	}{
		{
			"Invalid config",
			func() *config.Config {
				return nil
			},
			0,
			true,
		},
		{
			"Tx count",
			func() *config.Config {
				cfg, _ := config.ReadFromFile()
				return cfg
			},
			1,
			false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := GetTxCount(testCase.before())
			if err != nil {
				t.Fail()
			}
			if res.Result == 0 && !testCase.isError {
				t.Fail()
			}
		})
	}
}
