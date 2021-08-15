# go-weather

https://openweathermap.org/api

sample api

## Usage
```go
result := weather.GetWeather("${appId}", "${latitude}", "${longitude}")

fmt.Printf("result: %+v\n", result)
```

If the latitude and longitude are blank, the weather in Tokyo will be the default.
