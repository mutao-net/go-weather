package main

import (
	"fmt"

	"github.com/mutao-net/go-weather/weather"
)

func main() {
	result := weather.GetWeather("XXXXXX", "35.4660694", "139.6226196")

	// fmt.Printf("result: %+v\n", result)
	fmt.Println(result.Current.Feelslike)
	for _, value := range result.Current.Weather {
		fmt.Println(value.Description)
	}
}
