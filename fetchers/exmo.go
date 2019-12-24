package fetchers

import (
	"fmt"
	"net/http"
	"strings"
)

const exmoApi = "https://api.exmo.com/v1/ticker/"

type exmoFetcher struct {
	client *http.Client
}

type exmoResponse map[string]map[string]interface{}

func newExmoFetcher() exmoFetcher {
	return exmoFetcher{
		&http.Client{
			Timeout: defaultTimeout,
		},
	}
}

func (f exmoFetcher) fetch(symbols []string) (map[string]string, error) {
	resp, err := f.client.Get(exmoApi)
	if err != nil {
		return nil, fmt.Errorf("Fetch %s rates: %w", Exmo, err)
	}
	var allRatesInfo exmoResponse
	err = parseHttp(resp, &allRatesInfo)
	if err != nil {
		return nil, fmt.Errorf("Parse %s result: %w", Exmo, err)
	}

	allRates := make(map[string]string)
	for symbol, info := range allRatesInfo {
		allRates[symbol] = info["last_trade"].(string)
	}

	ratesNeeded := make(map[string]string)
	for _, symbol := range symbols {
		converted := strings.ReplaceAll(symbol, "/", "_")
		if allRates[converted] != "" {
			ratesNeeded[symbol] = allRates[converted]
		}
	}

	if len(ratesNeeded) == 0 {
		return nil, fmt.Errorf("Empty rates list for %s", Exmo)
	}
	return ratesNeeded, nil
}
