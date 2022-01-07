package temporalweather

import (
	"fmt"

	"github.com/helvellyn-io/weather/src/api"
)

func GetHumidity() (int, error) {
	var e error
	weather := new(api.Weather)
	api.GetJson("https://api.openweathermap.org/data/2.5/weather?q=Boulder&appid=050224087da57dabdc13099b40e747e0&units=imperial", weather)
	fmt.Printf("Location: %v Temp: %v F \n", weather.Name, weather.Main.Humidity)

	return weather.Main.Humidity, e

}
