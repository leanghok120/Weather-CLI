package main

import "github.com/leanghok120/moon/internal/weather"

func main() {
	flags := weather.ParseFlags()

	if flags.ShowVersion {
		weather.DisplayVersion(*flags)
	} else if flags.LongInfo {
		weather.DisplayLong(*flags)
	} else {
		weather.DisplayShort(*flags)
	}
}
