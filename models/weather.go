package models

type Location struct {
	Name string `json:"name"`
}

type Condition struct {
	Text string `json:"text"`
}

type CurrentWeather struct {
	Temperature float64   `json:"temp_c"`
	Condition   Condition `json:"condition"`
}

type WeatherData struct {
	Location Location       `json:"location"`
	Current  CurrentWeather `json:"current"`
}
