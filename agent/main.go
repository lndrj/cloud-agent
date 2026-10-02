package main

import (
	"agent/client"
	"agent/metrics"
	"log"
	"os"
	"time"
)

func main() {
	collectorURL := "http://localhost:8080"

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	for {
		metrics, err := metrics.CollectMetrics()
		if err != nil {
			log.Println("failed to collect metrics:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		if err := client.SendMetrics(
			collectorURL,
			hostname,
			metrics,
		); err != nil {
			log.Println("failed to send metrics:", err)
		}

		time.Sleep(5 * time.Second)
	}
}
