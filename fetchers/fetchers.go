package fetchers

import (
	"fmt"

	. "github.com/Croohand/crypto-monitor/markets"
	"github.com/Croohand/crypto-monitor/types"

	"github.com/Croohand/crypto-monitor/fetchers/binance"
	"github.com/Croohand/crypto-monitor/fetchers/exmo"
)

func GetAvailableMarkets() []Market {
	markets := make([]Market, 0, len(fetchers))
	for market := range fetchers {
		markets = append(markets, market)
	}
	return markets
}

func FetchRates(market Market, rates []types.Rate) ([]*types.RateInfo, error) {
	fetcher := fetchers[market]
	if fetcher == nil {
		return nil, fmt.Errorf("No fetcher available for market %s", market)
	}
	return fetcher.Fetch(rates)
}

type fetcher interface {
	Fetch(rates []types.Rate) ([]*types.RateInfo, error)
}

var fetchers = map[Market]fetcher{
	Binance: binance.New(),
	Exmo:    exmo.New(),
}
