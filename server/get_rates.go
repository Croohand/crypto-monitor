package server

import (
	"net/http"

	"github.com/Croohand/crypto-monitor/lib/fetchers"
	"github.com/Croohand/crypto-monitor/lib/helpers"
	"github.com/Croohand/crypto-monitor/lib/ratesbase"
)

func getRatesHandler(w http.ResponseWriter, r *http.Request) {
	markets := fetchers.GetAvailableMarkets()
	response := []map[string]string{}
	for _, market := range markets {
		for _, rate := range observedRates {
			rateInfo, err := ratesbase.Get(market, rate)
			if err != nil {
				continue
			}

			response = append(response, map[string]string{
				"pair":     rate.Symbol(),
				"exchange": string(market),
				"rate":     rateInfo.Price,
				"updated":  rateInfo.Updated.Format("2006-01-02 15:04:05.000"),
			})
		}
	}

	helpers.WriteHttp(w, response)
}
