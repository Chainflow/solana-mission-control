package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/Chainflow/solana-mission-control/config"
	"github.com/Chainflow/solana-mission-control/exporter"
	"github.com/Chainflow/solana-mission-control/monitor"
)

func main() {
	cfg, err := config.ReadFromFile() // Read config file
	if err != nil {
		log.Fatal(err)
	}

	collector := exporter.NewSolanaCollector(cfg)

	go collector.WatchSlots(cfg)

	// Calling command based alerting
	go func() {
		for {
			monitor.TelegramAlerting(cfg)
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			monitor.SkipRateAlerts(cfg)
			time.Sleep(60 * time.Second)
		}
	}()

	prometheus.MustRegister(collector)
	http.Handle("/metrics", promhttp.Handler()) // exported metrics can be seen in /metrics
	err = http.ListenAndServe(fmt.Sprintf("%s", cfg.Prometheus.ListenAddress), nil)
	if err != nil {
		log.Printf("Error while listening on server : %v", err)
	}
}
