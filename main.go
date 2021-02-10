package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/exporter"
)

func main() {
	// flag.Parse()

	// if *rpcAddr == "" {
	// 	klog.Fatal("Please specify -rpcURI")
	// }

	cfg, err := config.ReadFromFile()
	if err != nil {
		log.Fatal(err)
	}

	// log.Printf("cfg..", cfg)

	collector := exporter.NewSolanaCollector(cfg)

	go collector.WatchSlots(cfg)

	prometheus.MustRegister(collector)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)

	// klog.Infof("listening on %s", *addr)
	// klog.Fatal(http.ListenAndServe(*addr, nil))
}

// func recordMetrics() {
// 	go func() {
// 		for {
// 			opsProcessed.Inc()
// 			time.Sleep(2 * time.Second)
// 		}
// 	}()
// }

// var (
// 	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
// 		Name: "myapp_processed_ops_total",
// 		Help: "The total number of processed events",
// 	})
// )

// func main() {
// 	recordMetrics()

// 	http.Handle("/metrics", promhttp.Handler())
// 	http.ListenAndServe(":2112", nil)
// }
