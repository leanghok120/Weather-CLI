package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/leanghok120/moon/internal/weather"
)

func main() {
	location := weather.GetLocation()

	// Loads api key from .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Weather api stuff
	api_key := os.Getenv("API_KEY")
	const BASE_URL = "http://api.openweathermap.org/data/2.5/weather?"
	metric_unit := "metric"
	complete_url := BASE_URL + "appid=" + api_key + "&q=" + location + "&units=" + metric_unit

	weather := weather.GetWeather(complete_url)

	fmt.Printf("Location: %s\n", weather.Name)
	fmt.Printf("Temperature: %.0f°C\n", weather.Main.Temp)
	fmt.Printf("Humidity: %d\n", weather.Main.Humidity)
	fmt.Printf("Description: %s\n", weather.Weather[0].Description)
	fmt.Printf("Wind Speed: %.2f\n", weather.Wind.Speed)
}
