package fetchers

import "testing"

func TestExmoFetch(t *testing.T) {
	_, err := newExmoFetcher().fetch([]string{"ETH/BTC"})
	if err != nil {
		t.Error(err)
	}
}
