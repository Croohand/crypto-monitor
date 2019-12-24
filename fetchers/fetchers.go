package fetchers

import (
	"fmt"
	"time"
)

func GetAvailableMarkets() []Market {
	markets := make([]Market, 0, len(fetchers))
	for market := range fetchers {
		markets = append(markets, market)
	}
	return markets
}

func FetchRates(market Market, symbols []string) (map[string]string, error) {
	fetcher := fetchers[market]
	if fetcher == nil {
		return nil, fmt.Errorf("No fetcher available for market %s", market)
	}
	return fetcher.fetch(symbols)
}

const defaultTimeout = time.Duration(5 * time.Second)

type fetcher interface {
	fetch(symbols []string) (map[string]string, error)
}

var fetchers = map[Market]fetcher{
	Binance: newBinanceFetcher(),
	Exmo:    newExmoFetcher(),
}
