package darksky

import (
	"github.com/shawntoffel/darksky"
)

type Client struct {
	DarkSky darksky.DarkSky
}

type Interface interface {
	GetWeather(lat darksky.Measurement, lng darksky.Measurement, timestamp darksky.Timestamp) (*darksky.DataPoint, error)
}

// Create creates the new db server connection
func Create() Client {
	return Client{DarkSky: darksky.New("aac04c968b56377b27994b8a274ee502")}
}

func (c Client) GetWeather(lat darksky.Measurement, lng darksky.Measurement, timestamp darksky.Timestamp) (*darksky.DataPoint, error) {
	request := darksky.ForecastRequest{
		Latitude:  lat,       // 48.525300
		Longitude: lng,       // -123.389033
		Time:      timestamp, // 1550390399
		Options: darksky.ForecastRequestOptions{
			Exclude: "hourly,daily",
			Units:   "ca",
		},
	}
	resp, err := c.DarkSky.Forecast(request)
	return resp.Currently, err
}
