package exmo

import (
	"testing"

	"github.com/Croohand/crypto-monitor/types"
)

func TestExmoFetch(t *testing.T) {
	_, err := New().Fetch(types.ParseRates([]string{"ETH/BTC"}))
	if err != nil {
		t.Error(err)
	}
}