// external/weather.go
package external

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"signage-system/models"
)

// FetchWeather fetches the current weather data for a given location using the WeatherAPI.
// It returns the weather data in a structured format defined by the models.WeatherData type.
func FetchWeather(location string) (models.WeatherData, error) {
	// Get the API key from the environment variable
	key := os.Getenv("WEATHER_API_KEY")
	if key == "" {
		return models.WeatherData{}, errors.New("weather API key is missing")
	}

	// Check if the location is provided
	if location == "" {
		return models.WeatherData{}, errors.New("location is missing")
	}

	// Build the request URL using the location and API key
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", key, location)

	// Make the GET request to the WeatherAPI
	resp, err := http.Get(url)
	if err != nil {
		return models.WeatherData{}, fmt.Errorf("failed to fetch weather data: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		return models.WeatherData{}, fmt.Errorf("weather API returned status: %s", resp.Status)
	}

	// Decode the JSON response into the WeatherData struct
	var data models.WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return models.WeatherData{}, fmt.Errorf("failed to decode weather data: %v", err)
	}

	// Check if the essential fields in the response are not null or empty
	if data.Location.Name == "" || data.Current.Temperature == 0 {
		return models.WeatherData{}, errors.New("received incomplete weather data")
	}

	return data, nil
}
