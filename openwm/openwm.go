package openwm

// * The main "Logic" * //

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

const (
	weatherEP     = "https://api.openweathermap.org/data/2.5/weather"
	geoDirectEP   = "https://api.openweathermap.org/geo/1.0/direct"
	geoZipEP      = "https://api.openweathermap.org/geo/1.0/zip"
	locationLimit = "5" // TODO: make this configurable if it makes sense...
)

func NewOpenWM(apiKey string) *openWM {
	return &openWM{
		apiKey: apiKey,
	}
}

type openWM struct {
	apiKey string
}

func (wm *openWM) GetLocationsByCity(city string) (*Locations, error) {

	params := url.Values{}
	params.Add("q", city)
	params.Add("limit", locationLimit)
	params.Add("appid", wm.apiKey)

	return GetAndDecode[Locations](geoDirectEP + "?" + params.Encode())
}

// Usage: GetLocationInfoByZip("ZIP,COUNTRY_CODE")
func (wm *openWM) GetLocationInfoByZip(zip string) (*LocationInfo, error) {

	params := url.Values{}
	params.Add("zip", zip)
	params.Add("limit", locationLimit)
	params.Add("appid", wm.apiKey)

	return GetAndDecode[LocationInfo](geoZipEP + "?" + params.Encode())
}

func (wm *openWM) GetWeatherByLatLon(lat string, lon string) (*WeatherData, error) {

	params := url.Values{}
	params.Add("lat", lat)
	params.Add("lon", lon)
	params.Add("appid", wm.apiKey)

	return GetAndDecode[WeatherData](weatherEP + "?" + params.Encode())
}

func GetAndDecode[T Locations | LocationInfo | WeatherData](url string) (*T, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Error: remote server response was not '200 OK', it was: " + resp.Status)
	}

	var target T
	err = json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		return nil, err
	}
	return &target, nil
}
