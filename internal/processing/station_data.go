package processing

import (
	"time"
)

// StationData represents the data for a bike station at a specific timestamp
type StationData struct {
	Timestamp         time.Time `json:"timestamp"`
	StationID         string    `json:"station_id"`
	StationName       string    `json:"station_name"`
	Longitude         float64   `json:"longitude"`
	Latitude          float64   `json:"latitude"`
	Capacity          int       `json:"capacity"`
	BikesAvailable    int       `json:"bikes_available"`
	EBikesAvailable   int       `json:"ebikes_available"`
	ScootersAvailable int       `json:"scooters_available"`
	DocksAvailable    int       `json:"docks_available"`
	IsRenting         bool      `json:"is_renting"`
	IsReturning       bool      `json:"is_returning"`
	LastReported      int64     `json:"last_reported"`
}

// StationLogs maps station IDs to their respective logs
var StationLogs = make(map[string][]StationData)
