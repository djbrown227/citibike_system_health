package main

import (
	"go_citibike/internal/api" // Adjust the import path to match your module
	"log"
	"math"
	"sort"
	"time"
)

type StationDistance struct {
	Station  api.StationInfo
	Distance float64
}

// Haversine formula to calculate distance between two latitude/longitude points in miles
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 3958.8 // Earth radius in miles
	lat1Rad := lat1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c // returns distance in miles
}

func findClosestStations(station api.StationInfo, stations []api.StationInfo) []StationDistance {
	var distances []StationDistance

	for _, otherStation := range stations {
		if station.StationID != otherStation.StationID {
			distance := haversine(station.Lat, station.Lon, otherStation.Lat, otherStation.Lon)
			distances = append(distances, StationDistance{Station: otherStation, Distance: distance})
		}
	}

	// Sort distances
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	// Return the 10 closest stations
	if len(distances) > 10 {
		return distances[:10]
	}
	return distances
}

func main() {
	// Set log flags to include date and time in the output
	log.SetFlags(0) // Disable log timestamps

	log.Println("Starting CitiBike Data Fetcher...")

	// Fetch static station information only once since it does not change often
	staticStations := api.FetchStationInfo()

	// Create a map to index static stations by StationID
	stationMap := make(map[string]api.StationInfo)
	for _, station := range staticStations {
		stationMap[station.StationID] = station
	}

	// Fetch and join data every 5 seconds
	for {
		// Fetch dynamic bike availability data
		dynamicStations := api.FetchData()

		// Join the data using StationID
		for _, dynamicStation := range dynamicStations {
			if staticStation, ok := stationMap[dynamicStation.StationID]; ok {
				// Calculate % Filled feature
				percentFilled := float64(dynamicStation.NumBikesAvailable) / float64(staticStation.Capacity) * 100

				// Print combined information with structured output
				log.Printf("\n------------------------")
				log.Printf("Timestamp: %s", time.Now().Format("2006-01-02 15:04:05"))
				log.Printf("Station: %s (ID: %s)", staticStation.Name, staticStation.StationID)
				log.Printf("Longitude: %f", staticStation.Lon) // Longitude
				log.Printf("Latitude: %f", staticStation.Lat)  // Latitude
				log.Printf("Capacity: %d", staticStation.Capacity)
				log.Printf("Bikes Available: %d", dynamicStation.NumBikesAvailable)
				log.Printf("EBikes Available: %d", dynamicStation.NumEBikesAvailable)
				log.Printf("Scooters Available: %d", dynamicStation.NumScootersAvailable)
				log.Printf("Scooters Unavailable: %d", dynamicStation.NumScootersUnavailable)
				log.Printf("Bikes Disabled: %d", dynamicStation.NumBikesDisabled)
				log.Printf("Docks Available: %d", dynamicStation.NumDocksAvailable)
				log.Printf("Is Renting: %d", dynamicStation.IsRenting)
				log.Printf("Is Returning: %d", dynamicStation.IsReturning)
				log.Printf("Is Installed: %d", dynamicStation.IsInstalled)
				log.Printf("Last Reported: %d", dynamicStation.LastReported)

				// Print the new feature
				log.Printf("%% Filled: %.2f%%", percentFilled)

				// Find and print the 10 closest stations
				closestStations := findClosestStations(staticStation, staticStations)
				log.Println("Closest Stations:")
				for _, stationDistance := range closestStations {
					log.Printf(" - %s (ID: %s) at Distance: %.2f miles", stationDistance.Station.Name, stationDistance.Station.StationID, stationDistance.Distance)
				}

				log.Printf("------------------------\n")
			}
		}

		// Sleep for 5 seconds before the next fetch
		time.Sleep(5 * time.Second)
	}
}
