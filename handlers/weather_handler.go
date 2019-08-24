package handlers

import (
	"github.com/bolinkd/time-trial/middleware"
	"github.com/gin-gonic/gin"
)

func AddWeatherToTimeTrial(context *gin.Context) {
	// database := middleware.GetDatabase(context)
	darkskyClient := middleware.GetDarkSky(context)
	resp, err := darkskyClient.GetWeather(48.525300, -123.389033, 1550390399)
	if err != nil {
		UnexpectedError(context, err)
		return
	}
	Ok(context, resp)
}
