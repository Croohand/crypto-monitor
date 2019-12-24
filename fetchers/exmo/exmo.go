package exmo

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Croohand/crypto-monitor/helpers"
	"github.com/Croohand/crypto-monitor/markets"
	"github.com/Croohand/crypto-monitor/types"
)

const apiLink = "https://api.exmo.com/v1/ticker/"
const timeout = time.Duration(5 * time.Second)

type fetcher struct {
	client *http.Client
}

type response map[string]map[string]interface{}

type partialRateInfo struct {
	price   string
	updated time.Time
}

func New() fetcher {
	return fetcher{
		&http.Client{
			Timeout: timeout,
		},
	}
}

func (f fetcher) Fetch(rates []types.Rate) ([]*types.RateInfo, error) {
	resp, err := f.client.Get(apiLink)
	if err != nil {
		return nil, fmt.Errorf("Fetch %s rates: %w", markets.Exmo, err)
	}
	var allRatesInfo response
	err = helpers.ParseHttp(resp, &allRatesInfo)
	if err != nil {
		return nil, fmt.Errorf("Parse %s result: %w", markets.Exmo, err)
	}

	allRates := make(map[string]partialRateInfo)
	for symbol, info := range allRatesInfo {
		allRates[symbol] = partialRateInfo{
			info["last_trade"].(string),
			time.Unix(int64(info["updated"].(float64)), 0),
		}
	}

	ratesNeeded := []*types.RateInfo{}
	for _, rate := range rates {
		symbol := rate.From + "_" + rate.To
		if _, ok := allRates[symbol]; ok {
			rateInfo := types.NewRateInfo(
				rate,
				allRates[symbol].price,
			)
			rateInfo.Updated = allRates[symbol].updated
			ratesNeeded = append(ratesNeeded, rateInfo)
		}
	}

	if len(ratesNeeded) == 0 {
		return nil, fmt.Errorf("Empty rates list for %s", markets.Exmo)
	}
	return ratesNeeded, nil
}
