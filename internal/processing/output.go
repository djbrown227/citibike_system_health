package processing

import (
	"go_citibike/internal/api"
	"log"
	"os"
	"time"
)

// File pointer for log file
var logFile *os.File

// SetupLogFile initializes the log file for storing station data
func SetupLogFile() {
	var err error
	logFile, err = os.OpenFile("station_data_1.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	log.SetOutput(logFile) // Redirect log output to the file
	log.Println("Logging started at", time.Now())
}

// LogStationData logs station details into the log file
func LogStationData(stationID string, name string, lon, lat float64, capacity, numBikesAvailable, numEBikesAvailable, numScootersAvailable, numDocksAvailable int, isRenting, isReturning bool, lastReported int64) {
	log.Printf(
		"Station ID: %s, Name: %s, Location: (%.5f, %.5f), Capacity: %d, Bikes Available: %d, EBikes Available: %d, Scooters Available: %d, Docks Available: %d, Renting: %t, Returning: %t, Last Reported: %v\n",
		stationID, name, lon, lat, capacity, numBikesAvailable, numEBikesAvailable, numScootersAvailable, numDocksAvailable, isRenting, isReturning, time.Unix(lastReported, 0),
	)
}

// CloseLogFile closes the log file when the program exits
func CloseLogFile() {
	if logFile != nil {
		log.Println("Closing log file at", time.Now())
		logFile.Close()
	}
}

// PrintStationDetails prints the details of a single station to the console.
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
	log.Printf("Is Renting: %t", dynamicStation.IsRenting != 0)     // Log as bool
	log.Printf("Is Returning: %t", dynamicStation.IsReturning != 0) // Log as bool
	log.Printf("Is Installed: %d", dynamicStation.IsInstalled)
	log.Printf("Last Reported: %v", time.Unix(dynamicStation.LastReported, 0)) // Convert timestamp to readable format
	log.Printf("%% Filled: %.2f%%", percentFilled)
	log.Printf("%% Empty: %.2f%%", percentEmpty)
	log.Printf("------------------------\n")
}

// PrintClosestStations prints the details of the 10 closest stations to the console.
func PrintClosestStations(closestStations []StationDistance) {
	log.Println("Closest Stations:")
	for _, stationDistance := range closestStations {
		log.Printf(" - %s (ID: %s) at Distance: %.2f miles", stationDistance.Station.Name, stationDistance.Station.StationID, stationDistance.Distance)
	}
}
