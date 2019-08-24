package middleware

import (
	"github.com/bolinkd/time-trial/darksky"
	"github.com/gin-gonic/gin"
)

var (
	DarkskyServerKey = "darksky"
)

func (s *Server) DarkSkyHandler(context *gin.Context) {
	if _, ok := context.Keys[DarkskyServerKey]; !ok {
		context.Set(DarkskyServerKey, s.Darksky)
	}
}

func GetDarkSky(context *gin.Context) darksky.Interface {
	return context.Keys[DarkskyServerKey].(darksky.Interface)
}
