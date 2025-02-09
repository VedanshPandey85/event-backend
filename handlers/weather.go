package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func FetchWeather(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	// date := r.URL.Query().Get("date")
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", location, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var weatherData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Ensure the response has the expected structure
	if _, ok := weatherData["weather"]; !ok {
		http.Error(w, "Invalid weather data format", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(weatherData)
}
