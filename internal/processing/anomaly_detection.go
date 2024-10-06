package processing

import (
	"fmt"
	"time"
)

// AnomalyDetection checks for suspicious patterns in station bike availability
func AnomalyDetection() {
	// Set thresholds
	bikeThreshold := 5
	timeWindow := 5 * time.Minute

	for stationID, logs := range StationLogs {
		for i := 1; i < len(logs); i++ {
			current := logs[i]
			previous := logs[i-1]

			// Check if the current and previous records are within the time window
			if current.Timestamp.Sub(previous.Timestamp) <= timeWindow {
				// Detect significant changes in bike availability
				if abs(current.BikesAvailable-previous.BikesAvailable) >= bikeThreshold {
					fmt.Printf("Anomaly detected at station %s (%s): Bikes available changed from %d to %d\n", current.StationName, stationID, previous.BikesAvailable, current.BikesAvailable)
				}
			}
		}
	}
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
