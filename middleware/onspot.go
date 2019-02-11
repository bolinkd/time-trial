package middleware

import (
	"github.com/businessinstincts/traxone/onspot"
	"github.com/gin-gonic/gin"
)

var (
	OnspotServerKey = "onspot-svc"
)

func (s *Server) OnspotServiceHandler(context *gin.Context) {
	if _, ok := context.Keys[OnspotServerKey]; !ok {
		context.Set(OnspotServerKey, s.Onspot)
	}
}

func GetOnspot(context *gin.Context) onspot.ClientInterface {
	return context.Keys[OnspotServerKey].(onspot.ClientInterface)
}
