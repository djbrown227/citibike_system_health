package main

import (
	"go_citibike/internal/api"
	"go_citibike/internal/processing"
	"log"
	"time"
)

func main() {
	log.SetFlags(0) // Disable log timestamps
	log.Println("Starting CitiBike Data Fetcher...")

	// Fetch static station information once
	staticStations := api.FetchStationInfo()

	// Create a map to index static stations by StationID
	stationMap := processing.CreateStationMap(staticStations)

	// Fetch and join data every 5 seconds
	for {
		dynamicStations := api.FetchData()

		// Process each dynamic station
		for _, dynamicStation := range dynamicStations {
			if staticStation, ok := stationMap[dynamicStation.StationID]; ok {
				// Calculate features
				percentFilled := processing.CalculatePercentFilled(dynamicStation.NumBikesAvailable, staticStation.Capacity)
				percentEmpty := processing.CalculatePercentEmpty(percentFilled)

				// Print combined information with structured output
				processing.PrintStationDetails(staticStation, dynamicStation, percentFilled, percentEmpty)

				// Find and print the 10 closest stations
				closestStations := processing.FindClosestStations(staticStation, staticStations)
				processing.PrintClosestStations(closestStations)
			}
		}

		time.Sleep(5 * time.Second)
	}
}
