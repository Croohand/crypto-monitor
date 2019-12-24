package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/Croohand/crypto-monitor/types"
)

type ServiceConfig struct {
	ListenAddr      string
	DbPath          string
	RatesUpdateFreq int
	ObservedSymbols []string
}

func (cfg ServiceConfig) Check() {
	if cfg.DbPath == "" {
		log.Fatal("Empty DbPath in config")
	}

	if cfg.ListenAddr == "" {
		log.Fatal("Empty ListenAddr in config")
	}

	if cfg.RatesUpdateFreq == 0 {
		log.Fatal("Empty RatesUpdateFreq in config")
	}

	if len(cfg.ObservedSymbols) == 0 {
		log.Fatal("Empty ObservedSymbols in config")
	}

	for _, symbol := range cfg.ObservedSymbols {
		if !types.CanParse(symbol) {
			log.Printf("Incorrect symbol %s in ObservedSymbols in config\n", symbol)
		}
	}
}

func LoadConfig(cfgPath string) ServiceConfig {
	cfgFile, err := os.Open(cfgPath)
	if err != nil {
		log.Fatalf("Opening config file: %v", err)
	}
	defer cfgFile.Close()

	data, err := ioutil.ReadAll(cfgFile)
	if err != nil {
		log.Fatalf("Reading config file: %v", err)
	}

	var cfg ServiceConfig
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("Parsing config file: %v", err)
	}

	return cfg
}
