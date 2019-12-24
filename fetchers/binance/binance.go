package binance

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Croohand/crypto-monitor/helpers"
	"github.com/Croohand/crypto-monitor/markets"
	"github.com/Croohand/crypto-monitor/types"
)

const apiLink = "https://api.binance.com/api/v3/ticker/price"
const timeout = time.Duration(5 * time.Second)

type fetcher struct {
	client *http.Client
}

type response []map[string]string

func New() fetcher {
	return fetcher{
		&http.Client{
			Timeout: timeout,
		},
	}
}

func (f fetcher) Fetch(rates []types.Rate) ([]types.RateInfo, error) {
	resp, err := f.client.Get(apiLink)
	if err != nil {
		return nil, fmt.Errorf("Fetch %s rates: %w", markets.Binance, err)
	}
	var allRatesList response
	err = helpers.ParseHttp(resp, &allRatesList)
	if err != nil {
		return nil, fmt.Errorf("Parse %s result: %w", markets.Binance, err)
	}

	allRates := make(map[string]string)
	for _, rate := range allRatesList {
		allRates[rate["symbol"]] = rate["price"]
	}

	ratesNeeded := []types.RateInfo{}
	for _, rate := range rates {
		symbol := rate.From + rate.To
		if allRates[symbol] != "" {
			ratesNeeded = append(ratesNeeded, types.NewRateInfo(
				rate,
				allRates[symbol],
			))
		}
	}

	if len(ratesNeeded) == 0 {
		return nil, fmt.Errorf("Empty rates list for %s", markets.Binance)
	}
	return ratesNeeded, nil
}
