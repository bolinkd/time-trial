package middleware

import (
	"github.com/gin-gonic/gin"
)

var (
	AppNexusServerKey = "appnexus-svc"
)

func (s *Server) AppNexusServiceHandler(context *gin.Context) {
	if _, present := context.Get(AppNexusServerKey); !present {
		context.Set(AppNexusServerKey, s.AppNexus)
	}
}
