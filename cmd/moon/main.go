package main

import "github.com/leanghok120/moon/internal/weather"

func main() {
	flags := weather.ParseFlags()

	if flags.LongInfo == true {
		weather.DisplayLong(*flags)
	} else {
		weather.DisplayShort(*flags)
	}
}
