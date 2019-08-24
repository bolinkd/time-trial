package middleware

import (
	"github.com/bolinkd/time-trial/socket"
	"github.com/gin-gonic/gin"
)

var (
	SocketServerKey = "socket-svc"
)

func (s *Server) SocketServiceHandler(context *gin.Context) {
	if _, ok := context.Keys[SocketServerKey]; !ok {
		context.Set(SocketServerKey, s.Socket)
	}
}

func GetSocket(context *gin.Context) socket.ClientInterface {
	return context.Keys[SocketServerKey].(socket.ClientInterface)
}
