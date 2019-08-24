package middleware

import (
	"github.com/bolinkd/time-trial/service"
	"github.com/gin-gonic/gin"
)

var (
	ServicesServerKey = "services"
)

func (s *Server) ServicesHandler(context *gin.Context) {
	if _, ok := context.Keys[ServicesServerKey]; !ok {
		context.Set(ServicesServerKey, s.Services)
	}
}

func GetServices(context *gin.Context) service.ServicesInterface {
	return context.Keys[ServicesServerKey].(service.ServicesInterface)
}
