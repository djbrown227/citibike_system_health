package processing

import (
	"go_citibike/internal/api"
	"math"
	"sort"
)

// StationDistance is a helper struct to store station info and distance
type StationDistance struct {
	Station  api.StationInfo
	Distance float64
}

// Haversine calculates the distance between two latitude/longitude points in miles.
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 3958.8 // Earth radius in miles
	lat1Rad := lat1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

// FindClosestStations finds the 10 closest stations to the input station.
func FindClosestStations(station api.StationInfo, stations []api.StationInfo) []StationDistance {
	var distances []StationDistance

	for _, otherStation := range stations {
		if station.StationID != otherStation.StationID {
			distance := haversine(station.Lat, station.Lon, otherStation.Lat, otherStation.Lon)
			distances = append(distances, StationDistance{Station: otherStation, Distance: distance})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	if len(distances) > 10 {
		return distances[:10]
	}
	return distances
}
