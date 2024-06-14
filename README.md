# moon

A **simple** CLI built in Go that fetches weather information using the OpenWeather API.

## Features

- Fetches current weather data for a specified city.
- Displays temperature, humidity, weather description, and wind speed.

## Prerequisites

- An API key from [OpenWeather](https://openweathermap.org/api)

## Installation

1. Clone the repository:
    ```bash
    git clone git@github.com:leanghok120/moon.git
    ```
2. Navigate to the project directory:
    ```bash
    cd moon/cmd/moon
    ```
3. Edit the main.go and add your OpenWeather api key:
    ```go
    api_key = "your_api_key_here"
    ```

4. Remove this part of the main.go:
    ```go
    // Loads api key from .env
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }
    ```

5. Install moon
    ```bash
    go install .
    ```

## Usage

Fetches weather info about city

```bash
moon london
```

Fetches weather info about country

```bash
moon netherlands
```



## Example

```bash
moon london
Location: London
Temperature: 15 celsius
Humidity: 82%
Description: light rain
Wind Speed: 4.1 m/s
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## Contact

For questions or support, please open an issue on the GitHub repository.
