package server

import (
	"time"

	"github.com/Croohand/crypto-monitor/lib/fetchers"
	"github.com/Croohand/crypto-monitor/lib/markets"
	"github.com/Croohand/crypto-monitor/lib/ratesbase"
	"github.com/Croohand/crypto-monitor/lib/types"
)

type fetchResult struct {
	market markets.Market
	rates  types.MarketRates
}

func fetchRatesInto(marketRates chan<- fetchResult) {
	markets := fetchers.GetAvailableMarkets()

	// Set goroutines (and outgoing HTTP requests as well) limit
	pool := make(chan bool, 5)

	done := make(chan bool, len(markets))
	for _, market := range markets {
		pool <- true
		market := market
		go func() {
			rates, err := fetchers.FetchRates(market, observedRates)
			if err == nil {
				marketRates <- fetchResult{market, rates}
			}
			done <- true
			<-pool
		}()
	}

	for i := 0; i < len(markets); i++ {
		<-done
	}

	close(marketRates)
}

func monitorRates(interval int) {
	for {
		marketRates := make(chan fetchResult)
		go fetchRatesInto(marketRates)

		for result := range marketRates {
			for rate, rateInfo := range result.rates {
				ratesbase.Set(result.market, rate, rateInfo)
			}
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}
}
