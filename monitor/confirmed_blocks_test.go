package monitor_test

import (
	"testing"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/monitor"
)

func TestConfirmedBlocks(t *testing.T) {
	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config :", err)
	}
	res, err := monitor.GetConfirmedBlocks(70539000, 70539097, cfg)
	if err != nil {
		t.Error("Error while getting confirmed blocks in given range")
	}
	if res == nil {
		t.Error("Expexted non empty result, but got empty result: ", res)
	}
	if res != nil {
		t.Log("Got confirmed Blocks between given range : ", res)
	}
}

// var tests = []struct {
// 	rangeStart int64
// 	rangeEnd   int64
// 	want       []int64
// }{
// 	{70539000, 70539097, []int64{70539000, 70539001, 70539002}},
// 	{5, 10, []int64{5, 6, 7, 8, 9, 10}},
// }

// func TestConfirmedBlocks(t *testing.T) {

// 	cfg, err := config.ReadFromFile()
// 	if err != nil {
// 		t.Error("Error while reading config :", err)
// 	}
// 	for _, tt := range tests {
// 		testname := fmt.Sprintf("%d,%d", tt.rangeStart, tt.rangeEnd)
// 		t.Run(testname, func(t *testing.T) {
// 			res, err := monitor.GetConfirmedBlocks(tt.rangeStart, tt.rangeEnd, cfg)
// 			if err != nil {
// 				t.Error("Error while fetching confirmed Blocks")
// 			}
// 			if res != tt.want {

// 			}
// 		})
// 	}
// }
