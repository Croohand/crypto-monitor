package server

import (
	"testing"

	"github.com/Croohand/crypto-monitor/lib/types"
)

func TestFetchNonEmpty(t *testing.T) {
	observedRates = []types.Rate{types.Rate{"ETH", "BTC"}}
	marketRates := make(chan fetchResult)
	go fetchRatesInto(marketRates)
	any := false
	for _ = range marketRates {
		any = true
	}
	if !any {
		t.Error("Empty fetch results")
	}
}
