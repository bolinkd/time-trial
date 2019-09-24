package transport

import (
	"net/http"
	"time"
)

// Connection represents an end-point connection with transport
type Connection interface {
	GetMessage() (message string, err error)
	WriteMessage(message string) error
	Close() error
	PingParams() (interval, timeout time.Duration)
}

// Transport represents a connection transport
type Transport interface {
	Connect(url string) (conn Connection, err error)
	HandleConnection(w http.ResponseWriter, r *http.Request) (conn Connection, err error)
	Serve(w http.ResponseWriter, r *http.Request)
	SetSid(sid string, conn Connection)
}
