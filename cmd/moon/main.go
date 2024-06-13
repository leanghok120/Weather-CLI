package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Loads api key from .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Weather api stuff
	api_key := os.Getenv("API_KEY")
	const BASE_URL = "http://api.openweathermap.org/data/2.5/weather?"
	metric_unit := "metric"
	COMPLETE_URL := BASE_URL + "appid=" + api_key + "&q=Phnom%20Penh" + "&units=" + metric_unit

	// Get weather info
	res, err := http.Get(COMPLETE_URL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather api not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
