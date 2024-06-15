package main

import (
	"fmt"

	"github.com/leanghok120/moon/internal/weather"
)

func main() {
	location := weather.GetLocation()

	weather := weather.GetWeather(location)

	fmt.Printf("Location: %s\n", weather.Name)
	fmt.Printf("Temperature: %.0f°C\n", weather.Main.Temp)
	fmt.Printf("Humidity: %d\n", weather.Main.Humidity)
	fmt.Printf("Description: %s\n", weather.Weather[0].Description)
	fmt.Printf("Wind Speed: %.2f\n", weather.Wind.Speed)
}
