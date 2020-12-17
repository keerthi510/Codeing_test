package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]

		data, err := query(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8080", nil)
}

func query(city string) (weatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=86b6148fc81bc8f9d481f973379ded00&q=" + city)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	return d, nil
}

type weatherData struct {
	Name string `json:"name"`

	Main struct {
		Kelvin   float64 `json:"temp"`
		Pressure float64 `json:"pressure"`
		Humidity float64 `json:"humidity"`
	} `json:"main"`

	Coord struct {
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lon"`
	} `json:"coord"`

	Sys struct {
		Sunrise float64 `json:"sunrise"`
		Sunset  float64 `json:"sunset"`
	} `json:"sys"`

	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`

	Wind struct {
		Speed  float64 `json:"speed"`
		Degree float64 `json:"deg"`
	} `json:"wind"`
}
