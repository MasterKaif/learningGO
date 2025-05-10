package main

import (
	"fmt";
	"encoding/json";
	"net/http";
	"os";
)

const apiKey = "889f91077f44dca4f3db5c8e2b52a2c5";

type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func main () {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run weather.go <city>")
		return
	}

	city := os.Args[1]
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to fetch weather Data:", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("City not found or API error:", resp.Status)
		fmt.Println("Response:", resp)
		return
	}

	var result WeatherResponse

	if err := json.NewDecoder((resp.Body)).Decode((&result)); err != nil {
		fmt.Println("Failed to decode JSON:", err)
		return
	}

	fmt.Printf("City: %s\n", result.Name)
	fmt.Printf("Temperature: %.2f°C\n", result.Main.Temp)
	fmt.Printf("Weather: %s\n", result.Weather[0].Description)

}