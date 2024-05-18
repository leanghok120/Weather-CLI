# Weather-CLI

A command-line interface (CLI) application built in Python that fetches weather information using the OpenWeather API.

## Features

- Fetches current weather data for a specified city.
- Displays temperature, humidity, weather description, and wind speed.

## Prerequisites

- Python 3.x
- An API key from [OpenWeather](https://openweathermap.org/api)

## Installation

1. Clone the repository:
    ```bash
    git clone git@github.com:leanghok120/Weather-CLI.git
    ```
2. Navigate to the project directory:
    ```bash
    cd Weather-CLI
    ```
3. Install required dependencies:
    ```bash
    pip install requests
    ```

## Usage

1. Open `main.py` and add your OpenWeather API key:
    ```python
    API_KEY = 'your_api_key_here'
    ```
2. Run the script:
    ```bash
    python main.py {city_name}
    ```

## Example

```bash
$ python main.py London
Temperature: 15 celsius
Humidity: 82%
Description: light rain
Wind Speed: 4.1 m/s
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## Contact

For questions or support, please open an issue on the GitHub repository.
