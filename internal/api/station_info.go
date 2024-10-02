package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// StationInfo represents the structure of static station information
type StationInfo struct {
	StationID  string  `json:"station_id"`
	Name       string  `json:"name"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	Capacity   int     `json:"capacity"`
	RegionID   string  `json:"region_id"`
	ShortName  string  `json:"short_name"`
	RentalUris struct {
		Android string `json:"android"`
		Ios     string `json:"ios"`
	} `json:"rental_uris"`
}

type StationInfoData struct {
	Stations []StationInfo `json:"stations"`
}

type StationInfoApiResponse struct {
	Data StationInfoData `json:"data"`
}

// FetchStationInfo retrieves the static station information
func FetchStationInfo() []StationInfo {
	url := "https://gbfs.lyft.com/gbfs/2.3/bkn/en/station_information.json"

	response, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching station information: %v", err)
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

	var stationInfoApiResponse StationInfoApiResponse
	if err := json.Unmarshal(body, &stationInfoApiResponse); err != nil {
		log.Printf("Error unmarshalling station information JSON: %v", err)
		return nil
	}

	return stationInfoApiResponse.Data.Stations
}
