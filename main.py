import sys
import requests
from dotenv import load_dotenv, dotenv_values 
import os
import json

load_dotenv()

def main(argv):
    CITY_NAME = argv[1]

    # Get weather info from open weather api
    API_KEY = os.getenv("OPEN_WEATHER_API")
    BASE_URL = "http://api.openweathermap.org/data/2.5/weather?"
    metric_unit = "metric"
    COMPLETE_URL = BASE_URL + "appid=" + API_KEY + "&q=" + CITY_NAME + "&units=" + metric_unit

    response = requests.get(COMPLETE_URL)
    response = response.json()

    # Check city and API Key
    if response["cod"] == "404":
        print("City not found.\nCheck out open weather api for more info", file=sys.stderr)
        sys.exit()
    elif response["cod"] == "401":
        print("Invalid api key, https://openweathermap.org/faq#error401", file=sys.stderr)
        sys.exit()

    # Print out temperature
    main = response["main"]
    temperature = main["temp"]
    print("Temperature: " + str(temperature) + " celsius")

    # Prints out humidity
    humidity = main["humidity"]
    print("Humidity: " + str(humidity) + "%")

    # Print out description
    weather = response["weather"]
    weather_dict = weather[0]
    description = weather_dict["description"]

    print("Description: " + str(description))

    # Print out wind speed
    wind = response["wind"]
    wind_speed = wind["speed"]

    print("Wind Speed: " + str(wind_speed) + " m/s")

if __name__ == "__main__":
    main(sys.argv)
