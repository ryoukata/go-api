package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

// 天気データAPIから取得したデータの保存型
type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

// すべての気象データAPIで同じ振る舞いにするためのインターフェース
type weatherProvider interface {
	temperature(city string) (float64, error) // in Kelvin, naturally
}

type openWeatherMap struct{}

type multiWeatherProvider []weatherProvider

type weatherUnderground struct {
	apiKey string
}

// GetWeatherByCity return echo.HandlerFunc
func GetWeatherByCity() echo.HandlerFunc {
	return func(c echo.Context) error {
		city := c.Param("city")
		data, err := query(city)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Server Error: "+err.Error())
		}

		return c.JSON(http.StatusOK, data)
	}
}

// GetWeather return echo.HandlerFunc
func GetWeather() echo.HandlerFunc {
	return func(c echo.Context) error {
		mw := multiWeatherProvider{
			openWeatherMap{},
			weatherUnderground{apiKey: "API_Key"},
		}

		begin := time.Now()
		city := c.Param("city")

		temp, err := mw.temperature(city)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Server Error: "+err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"city": city,
			"temp": temp,
			"took": time.Since(begin).String(),
		})
	}
}

// query return (weatherData, error)
func query(city string) (weatherData, error) { // APIから天気データを取得する関数
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&APPID=API_Key")
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var data weatherData

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return weatherData{}, err
	}

	return data, nil
}

// temperature return (float64, error)
func (w openWeatherMap) temperature(city string) (float64, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&APPID=API_Key")
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	var data struct {
		Main struct {
			Kelvin float64 `json:"temp"`
		} `json:"main"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	log.Printf("openWeatherMap: %s: %.2f", city, data.Main.Kelvin)
	return data.Main.Kelvin, nil
}

// temperature return (float64, error)
func (w weatherUnderground) temperature(city string) (float64, error) {
	resp, err := http.Get("http://api.wunderground.com/api/" + w.apiKey + "/conditions/q/" + city + ".json")
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	var data struct {
		Observation struct {
			Celsius float64 `json:"temp_c"`
		} `json:"current_observation"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	Kelvin := data.Observation.Celsius + 273.15

	log.Printf("weatherUnderground: %s: %.2f", city, Kelvin)
	return Kelvin, nil
}

// temperature return (float64, error)
func temperature(city string, providers ...weatherProvider) (float64, error) { // 上記の両temperatureにクエリし、平均気温を返す関数
	sum := 0.0

	for _, provider := range providers {
		k, err := provider.temperature(city)
		if err != nil {
			return 0, err
		}

		sum += k
	}

	return sum / float64(len(providers)), nil
}

func (w multiWeatherProvider) temperature(city string) (float64, error) {
	sum := 0.0

	for _, provider := range w {
		k, err := provider.temperature(city)
		if err != nil {
			return 0, err
		}

		sum += k
	}

	return sum / float64(len(w)), nil
}
