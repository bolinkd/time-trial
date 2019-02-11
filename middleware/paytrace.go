package middleware

import (
	"github.com/businessinstincts/traxone/paytrace"
	"github.com/gin-gonic/gin"
)

var (
	PayTraceServerKey = "paytrace-svc"
)

func (s *Server) PayTraceHandler(context *gin.Context) {
	if _, present := context.Get(PayTraceServerKey); !present {
		context.Set(PayTraceServerKey, s.PayTrace)
	}
}

func GetPayTrace(context *gin.Context) paytrace.ClientInterface {
	return context.Keys[PayTraceServerKey].(paytrace.ClientInterface)
}
