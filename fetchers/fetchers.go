package fetchers

import (
	"errors"
)

func GetMarkets() []string {
	markets := make([]string, 0, len(fetchers))
	for market := range fetchers {
		markets = append(markets, market)
	}
	return markets
}

func FetchRates(market string, symbols []string) (map[string]string, error) {
	fetcher := fetchers[market]
	if fetcher == nil {
		return nil, errors.New("No fetcher available for market " + market)
	}
	return fetcher.fetch(symbols)
}

type fetcher interface {
	fetch(symbols []string) (map[string]string, error)
}

var fetchers = map[string]fetcher{}
