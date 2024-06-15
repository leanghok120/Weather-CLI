package weather

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Weather struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp     float32 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float32 `json:"speed"`
	} `json:"wind"`
	Name string `json:"name"`
}

func GetWeather(location string) Weather {
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

	// Get weather info
	res, err := http.Get(complete_url)
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

	// Parse body into weather struct
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	return weather
}
