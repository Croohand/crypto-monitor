package server

import (
	"fmt"
	"net/http"

	"github.com/Croohand/crypto-monitor/config"
	"github.com/Croohand/crypto-monitor/lib/ratesbase"
	"github.com/Croohand/crypto-monitor/lib/types"
)

var observedRates []types.Rate

func Run(cfg config.ServiceConfig) error {
	err := ratesbase.Open(cfg.DbPath)
	if err != nil {
		return fmt.Errorf("Open rates database: %w", err)
	}
	defer ratesbase.Close()

	observedRates = types.ParseRates(cfg.ObservedSymbols)
	go monitorRates(cfg.RatesUpdateFreq)

	routes()

	return http.ListenAndServe(cfg.ListenAddr, nil)
}

func routes() {
	http.HandleFunc("/get-rates", getRatesHandler)
}
