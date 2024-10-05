package processing

import (
	"go_citibike/internal/api"
	"log"
	"time"
)

// PrintStationDetails prints the details of a single station.
func PrintStationDetails(staticStation api.StationInfo, dynamicStation api.Station, percentFilled, percentEmpty float64) {
	log.Printf("\n------------------------")
	log.Printf("Timestamp: %s", time.Now().Format("2006-01-02 15:04:05"))
	log.Printf("Station: %s (ID: %s)", staticStation.Name, staticStation.StationID)
	log.Printf("Longitude: %f", staticStation.Lon)
	log.Printf("Latitude: %f", staticStation.Lat)
	log.Printf("Region ID: %s", staticStation.RegionID)
	log.Printf("Capacity: %d", staticStation.Capacity)
	log.Printf("Bikes Available: %d", dynamicStation.NumBikesAvailable)
	log.Printf("EBikes Available: %d", dynamicStation.NumEBikesAvailable)
	log.Printf("Scooters Available: %d", dynamicStation.NumScootersAvailable)
	log.Printf("Bikes Disabled: %d", dynamicStation.NumBikesDisabled)
	log.Printf("Docks Available: %d", dynamicStation.NumDocksAvailable)
	log.Printf("Docks Disabled: %d", dynamicStation.NumDocksDisabled)
	log.Printf("Is Renting: %d", dynamicStation.IsRenting)
	log.Printf("Is Returning: %d", dynamicStation.IsReturning)
	log.Printf("Is Installed: %d", dynamicStation.IsInstalled)
	log.Printf("Last Reported: %d", dynamicStation.LastReported)
	log.Printf("%% Filled: %.2f%%", percentFilled)
	log.Printf("%% Empty: %.2f%%", percentEmpty)
	log.Printf("------------------------\n")
}

// PrintClosestStations prints the details of the 10 closest stations.
func PrintClosestStations(closestStations []StationDistance) {
	log.Println("Closest Stations:")
	for _, stationDistance := range closestStations {
		log.Printf(" - %s (ID: %s) at Distance: %.2f miles", stationDistance.Station.Name, stationDistance.Station.StationID, stationDistance.Distance)
	}
}
