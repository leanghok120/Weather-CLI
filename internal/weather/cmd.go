package weather

import (
	"fmt"
)

func DisplayShort() {
	location := GetLocation()
	weather := GetWeather(location)

	fmt.Printf("Location: %s\n", weather.Name)
	fmt.Printf("Temperature: %.0f°C\n", weather.Main.Temp)
	fmt.Printf("Humidity: %d\n", weather.Main.Humidity)
	fmt.Printf("Description: %s\n", weather.Weather[0].Description)
	fmt.Printf("Wind Speed: %.2f\n", weather.Wind.Speed)
}

func DisplayLong() {
	location := GetLocation()
	weather := GetWeather(location)

	fmt.Printf("Location: %s\n", weather.Name)
	fmt.Printf("Temperature: %.0f°C\n", weather.Main.Temp)
	fmt.Printf("Feels like: %.0f°C\n", weather.Main.Feels_like)
	fmt.Printf("Humidity: %d\n", weather.Main.Humidity)
	fmt.Printf("Description: %s\n", weather.Weather[0].Description)
	fmt.Printf("Wind Speed: %.2f\n", weather.Wind.Speed)
}
