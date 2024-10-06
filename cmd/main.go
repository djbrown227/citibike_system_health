package main

import (
	"flag"
	"go_citibike/internal/api"
	"go_citibike/internal/processing"
	"log"
	"math"
	"sort"
	"time"
)

// Struct to hold station change information
type StationChange struct {
	StationID         string
	Name              string
	TotalFilledChange float64
	TotalEmptyChange  float64
}

func main() {
	var durationSeconds int
	flag.IntVar(&durationSeconds, "duration", 600, "Duration to run the program in seconds (default is 6000 seconds)")
	flag.Parse()

	log.SetFlags(0) // Disable log timestamps

	// Setup log file
	processing.SetupLogFile()
	defer processing.CloseLogFile() // Ensure the log file is closed at program exit

	log.Println("Starting CitiBike Data Fetcher...")

	// Fetch static station information once
	staticStations := api.FetchStationInfo()
	stationMap := processing.CreateStationMap(staticStations)
	previousMetrics := make(map[string]struct {
		percentFilled float64
		percentEmpty  float64
	})
	totalFilledChanges := make(map[string]float64)
	totalEmptyChanges := make(map[string]float64)

	startTime := time.Now()

	for {
		if time.Since(startTime) >= time.Duration(durationSeconds)*time.Second {
			log.Println("Duration has elapsed, stopping the program.")
			break
		}

		dynamicStations := api.FetchData()
		for _, dynamicStation := range dynamicStations {
			if staticStation, ok := stationMap[dynamicStation.StationID]; ok {
				percentFilled := processing.CalculatePercentFilled(dynamicStation.NumBikesAvailable, staticStation.Capacity)

				// Debug: log the current filled percentage and station details
				log.Printf("Station %s (ID: %s): NumBikesAvailable: %d, Capacity: %d, PercentFilled: %.2f%%",
					staticStation.Name, dynamicStation.StationID, dynamicStation.NumBikesAvailable, staticStation.Capacity, percentFilled)

				// Calculate percentEmpty based on the filled percentage
				percentEmpty := 100 - percentFilled

				// Debug: log the current empty percentage
				log.Printf("Station %s (ID: %s): PercentEmpty: %.2f%%", staticStation.Name, dynamicStation.StationID, percentEmpty)

				if previous, exists := previousMetrics[staticStation.StationID]; exists {
					// Calculate absolute changes in percent filled and percent empty
					filledChange := math.Abs(percentFilled - previous.percentFilled)
					emptyChange := math.Abs(percentEmpty - previous.percentEmpty)

					// Accumulate the total changes
					totalFilledChanges[staticStation.StationID] += filledChange
					totalEmptyChanges[staticStation.StationID] += emptyChange

					// Debug: log the absolute change in filled and empty percentages
					log.Printf("Absolute change in %% Filled for %s (ID: %s): %.2f%%, Absolute change in %% Empty: %.2f%%",
						staticStation.Name, staticStation.StationID, filledChange, emptyChange)
				}

				// Update the previous metrics for the station
				previousMetrics[staticStation.StationID] = struct {
					percentFilled float64
					percentEmpty  float64
				}{
					percentFilled,
					percentEmpty,
				}

				processing.PrintStationDetails(staticStation, dynamicStation, percentFilled, percentEmpty)

				isRenting := dynamicStation.IsRenting != 0
				isReturning := dynamicStation.IsReturning != 0
				processing.LogStationData(
					dynamicStation.StationID,
					staticStation.Name,
					staticStation.Lon,
					staticStation.Lat,
					staticStation.Capacity,
					dynamicStation.NumBikesAvailable,
					dynamicStation.NumEBikesAvailable,
					dynamicStation.NumScootersAvailable,
					dynamicStation.NumDocksAvailable,
					isRenting,
					isReturning,
					dynamicStation.LastReported,
				)
			}
		}

		processing.AnomalyDetection()
		time.Sleep(10 * time.Second)
	}

	log.Println("\nSummary of Changes in % Filled and % Empty:")
	var stationChanges []StationChange
	for stationID := range totalFilledChanges {
		if staticStation, ok := stationMap[stationID]; ok {
			stationChanges = append(stationChanges, StationChange{
				StationID:         stationID,
				Name:              staticStation.Name,
				TotalFilledChange: totalFilledChanges[stationID],
				TotalEmptyChange:  totalEmptyChanges[stationID],
			})
		}
	}

	// Sort the stations by the total filled change
	sort.Slice(stationChanges, func(i, j int) bool {
		return stationChanges[i].TotalFilledChange > stationChanges[j].TotalFilledChange
	})

	// Log the summary of changes for each station
	for _, stationChange := range stationChanges {
		log.Printf("Station: %s (ID: %s), Total Filled Change: %.2f%%, Total Empty Change: %.2f%%",
			stationChange.Name,
			stationChange.StationID,
			stationChange.TotalFilledChange,
			stationChange.TotalEmptyChange)
	}
}
