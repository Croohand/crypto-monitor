package fetchers

import "testing"

func TestBinanceFetch(t *testing.T) {
	_, err := newBinanceFetcher().fetch([]string{"ETH/BTC"})
	if err != nil {
		t.Error(err)
	}
}
