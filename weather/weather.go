package weather

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const api = "https://api.openweathermap.org/data/2.5/onecall"

type ResponseResult struct {
	Lat            float64        `json:"lat"`
	Lon            float64        `json:"lon"`
	Timezone       string         `json:"timezone"`
	TimezoneOffset int            `json:"timezone_offset"`
	Current        CurrentEntitys `json:"current"`
}

type CurrentEntitys struct {
	Dt         int              `json:"dt"`
	Sunrise    int              `json:"sunrise"`
	Sunset     int              `json:"sunset"`
	Temp       float64          `json:"temp"`
	Feelslike  float64          `json:"feels_like"`
	Pressure   int              `json:"pressure"`
	Humidity   int              `json:"humidity"`
	DewPoint   float64          `json:"dew_point"`
	Visibility int              `json"visibility"`
	WindSpeed  float64          `json:"wind_speed"`
	WindDeg    int              `json:"wind_deg"`
	WindGust   float64          `json:"wind_gust"`
	Weather    []WeatherEntitys `json:"weather"`
}

type WeatherEntitys struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

func GetWeather(appId string, lat string, lon string) ResponseResult {
	req := setParams(appId, lat, lon)
	result := callAPI(req)
	return result
}

func setParams(appId string, lat string, lon string) string {
	//default Tokyo
	if len(lat) == 0 {
		lat = "35.681236"
	}
	if len(lon) == 0 {
		lon = "139.767125"
	}

	queries := []string{}
	params := map[string]string{
		"lat":   lat,
		"lon":   lon,
		"units": "metric",
		"lang":  "ja",
		"appid": appId}
	for k, v := range params {
		query := k + "=" + url.QueryEscape(v)
		queries = append(queries, query)
	}
	return api + "?" + strings.Join(queries, "&")
}

func callAPI(url string) ResponseResult {
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		log.Fatal("Get Http Error:", err)
	}
	var results ResponseResult
	json.NewDecoder(response.Body).Decode(&results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}
