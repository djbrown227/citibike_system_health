package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// VehicleType represents the available vehicle types and their counts
type VehicleType struct {
	Count         int    `json:"count"`
	VehicleTypeID string `json:"vehicle_type_id"`
}

// Station represents the dynamic station data
type Station struct {
	VehicleTypesAvailable  []VehicleType `json:"vehicle_types_available"`
	NumDocksDisabled       int           `json:"num_docks_disabled"`
	NumScootersUnavailable int           `json:"num_scooters_unavailable"`
	NumBikesAvailable      int           `json:"num_bikes_available"`
	LastReported           int64         `json:"last_reported"`
	NumDocksAvailable      int           `json:"num_docks_available"`
	NumEBikesAvailable     int           `json:"num_ebikes_available"`
	IsReturning            int           `json:"is_returning"`
	IsRenting              int           `json:"is_renting"`
	IsInstalled            int           `json:"is_installed"`
	NumBikesDisabled       int           `json:"num_bikes_disabled"`
	NumScootersAvailable   int           `json:"num_scooters_available"`
	StationID              string        `json:"station_id"`
}

type Data struct {
	Stations []Station `json:"stations"`
}

type ApiResponse struct {
	Data Data `json:"data"`
}

// FetchData retrieves the dynamic bike availability data from the Citi Bike API
func FetchData() []Station {
	url := "https://gbfs.citibikenyc.com/gbfs/2.3/en/station_status.json"

	response, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching data: %v", err)
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("Error: received status code %d", response.StatusCode)
		return nil
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil
	}

	var apiResponse ApiResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return nil
	}

	return apiResponse.Data.Stations
}
