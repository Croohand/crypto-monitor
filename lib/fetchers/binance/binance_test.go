package binance

import (
	"testing"

	. "github.com/Croohand/crypto-monitor/lib/types"
)

func TestBinanceFetch(t *testing.T) {
	_, err := New().Fetch(ParseRates([]string{"ETH/BTC"}))
	if err != nil {
		t.Error(err)
	}
}
