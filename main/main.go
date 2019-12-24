package main

import (
	"flag"
	"log"

	"github.com/Croohand/crypto-monitor/config"
	"github.com/Croohand/mapreduce/slave/server"
)

func main() {
	cfgPath := flag.String("config", "", "Path to config.json")

	flag.Parse()

	if *cfgPath == "" {
		log.Fatal("Empty config path")
	}

	cfg := config.LoadConfig(*cfgPath)
	cfg.Check()

	log.Printf("Starting service with config %#v\n", cfg)

	err := server.Run(cfg)
}
