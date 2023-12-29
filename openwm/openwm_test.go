package openwm

// * Some basic tests * //

import (
	"testing"
)

const (
	apiKeyTest = "ADD_KEY_FOR_TESTING_HERE"
)

func TestOpenWM(t *testing.T) {

	openWM := NewOpenWM(apiKeyTest)

	// Testing GetWeatherByLatLon: WeatherData
	wd, err := openWM.GetWeatherByLatLon("63.83847", "23.13066")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if wd.Name != "Kokkola" {
		t.Errorf("Expected Kokkola, but got %s", wd.Name)
	}

	// Testing GetLocationsByCity: Locations
	ld, err := openWM.GetLocationsByCity("Kokkola")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(*ld) > 3 {
		t.Errorf("Expected max 3 locations, but got %d", len(*ld))
	}
	for _, location := range *ld {
		if location.Name != "Kokkola" {
			t.Errorf("Expected Kokkola, but got %s", location.Name)
		}
	}

	// Testing GetLocationInfoByZip: LocationInfo
	li, err := openWM.GetLocationInfoByZip("67100,FI")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if li.Name != "Kokkola" {
		t.Errorf("Expected Kokkola, but got %s", li.Name)
	}
}

func TestGetAndDecode(t *testing.T) {

	// Test Locations
	locations, err := GetAndDecode[Locations]("https://api.openweathermap.org/geo/1.0/direct?q=Kokkola&limit=3&appid=" + apiKeyTest)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(*locations) > 3 {
		t.Errorf("Expected 3 locations at maximum, but got %d", len(*locations))
	}
	for _, location := range *locations {
		if location.Name != "Kokkola" {
			t.Errorf("Expected Kokkola, but got %s", location.Name)
		}
	}

	// Test LocationInfo
	locationInfo, err := GetAndDecode[LocationInfo]("https://api.openweathermap.org/geo/1.0/zip?zip=67100,FI&limit=3&appid=" + apiKeyTest)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if locationInfo.Name != "Kokkola" {
		t.Errorf("Expected Kokkola, but got %s", locationInfo.Name)
	}

	// Test WeatherData
	weatherData, err := GetAndDecode[WeatherData]("https://api.openweathermap.org/data/2.5/weather?lat=63.83847&lon=23.13066&appid=" + apiKeyTest)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if weatherData.Name != "Kokkola" {
		t.Errorf("Expected Kokkola, but got %s", weatherData.Name)
	}

	//Invalid api key Test
	_, err = GetAndDecode[WeatherData]("https://api.openweathermap.org/data/2.5/weather?lat=63.83847&lon=23.13066&appid=XXXX")
	if err == nil {
		t.Errorf("Error, wanted err got : %s", err)
	}

	//Incorrect url Test
	_, err = GetAndDecode[WeatherData]("http://google.com")
	if err == nil {
		t.Errorf("Error, wanted err got : %s", err)
	}

	//Invalid url Test
	_, err = GetAndDecode[WeatherData]("INVALID URL")
	if err == nil {
		t.Errorf("Error, wanted err got : %s", err)
	}
}
