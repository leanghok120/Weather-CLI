package weather

import (
	"flag"
	"fmt"
)

type Config struct {
	Location    string
	LongInfo    bool
	ShowVersion bool
}

func ParseFlags() *Config {
	var location string
	var longInfo bool
	var showVersion bool

	flag.StringVar(&location, "c", "london", "Location")
	flag.BoolVar(&longInfo, "L", false, "Show more information")
	flag.BoolVar(&showVersion, "v", false, "Show version")

	flag.Parse()

	return &Config{
		Location:    location,
		LongInfo:    longInfo,
		ShowVersion: showVersion,
	}
}

func DisplayShort(flags Config) {
	location := GetLocation(flags)
	weather := GetWeather(location)

	fmt.Printf("Location: %s\n", weather.Name)
	fmt.Printf("Temperature: %.0f°C\n", weather.Main.Temp)
	fmt.Printf("Humidity: %d\n", weather.Main.Humidity)
	fmt.Printf("Description: %s\n", weather.Weather[0].Description)
	fmt.Printf("Wind Speed: %.2f\n", weather.Wind.Speed)
}

func DisplayLong(flags Config) {
	location := GetLocation(flags)
	weather := GetWeather(location)

	fmt.Printf("Location: %s\n", weather.Name)
	fmt.Printf("Temperature: %.0f°C\n", weather.Main.Temp)
	fmt.Printf("Feels like: %.0f°C\n", weather.Main.Feels_like)
	fmt.Printf("Humidity: %d\n", weather.Main.Humidity)
	fmt.Printf("Description: %s\n", weather.Weather[0].Description)
	fmt.Printf("Wind Speed: %.2f\n", weather.Wind.Speed)
}
