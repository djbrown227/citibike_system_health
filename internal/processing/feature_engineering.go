package processing

import (
	"go_citibike/internal/api"
)

// CalculatePercentFilled calculates the percentage of bikes available.
func CalculatePercentFilled(numBikesAvailable, capacity int) float64 {
	if capacity == 0 {
		return 0
	}
	return float64(numBikesAvailable) / float64(capacity) * 100
}

// CalculatePercentEmpty calculates the percentage of empty docks.
func CalculatePercentEmpty(percentFilled float64) float64 {
	return 100 - percentFilled
}

// CreateStationMap indexes stations by their StationID for quick lookups.
func CreateStationMap(stations []api.StationInfo) map[string]api.StationInfo {
	stationMap := make(map[string]api.StationInfo)
	for _, station := range stations {
		stationMap[station.StationID] = station
	}
	return stationMap
}
