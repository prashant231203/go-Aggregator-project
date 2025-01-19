package main

import (
	"log"
	"time"

	"github.com/prashant231203/go-Aggregator-project/internal/database"
)

func startScraping(
	db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping on %v goroutiens every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {

	}
}
