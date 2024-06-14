package weather

import "os"

func GetLocation() string {
	location := "london"

	if len(os.Args) >= 2 {
		location = os.Args[1]
	}

	return location
}
