package fetchers

import (
	"fmt"
	"net/http"
	"strings"
)

const binanceApi = "https://api.binance.com/api/v3/ticker/price"

type binanceFetcher struct {
	client *http.Client
}

type binanceResponse []map[string]string

func newBinanceFetcher() binanceFetcher {
	return binanceFetcher{
		&http.Client{
			Timeout: defaultTimeout,
		},
	}
}

func (f binanceFetcher) fetch(symbols []string) (map[string]string, error) {
	resp, err := f.client.Get(binanceApi)
	if err != nil {
		return nil, fmt.Errorf("Fetch %s rates: %w", Binance, err)
	}
	var allRatesList binanceResponse
	err = parseHttp(resp, &allRatesList)
	if err != nil {
		return nil, fmt.Errorf("Parse %s result: %w", Binance, err)
	}

	allRates := make(map[string]string)
	for _, rate := range allRatesList {
		allRates[rate["symbol"]] = rate["price"]
	}

	ratesNeeded := make(map[string]string)
	for _, symbol := range symbols {
		converted := strings.ReplaceAll(symbol, "/", "")
		if allRates[converted] != "" {
			ratesNeeded[symbol] = allRates[converted]
		}
	}

	if len(ratesNeeded) == 0 {
		return nil, fmt.Errorf("Empty rates list for %s", Binance)
	}
	return ratesNeeded, nil
}
