package middleware

import (
	"errors"
	"fmt"
	"github.com/businessinstincts/traxone/appnexus"
	"github.com/businessinstincts/traxone/db"
	"github.com/businessinstincts/traxone/onspot"
	"github.com/businessinstincts/traxone/paytrace"
	"github.com/businessinstincts/traxone/socket"
	"time"
)

// Server structure
type Server struct {
	Database *db.Connection
	Onspot   *onspot.Client
	Socket   *socket.Client
	AppNexus *appnexus.Client
	PayTrace *paytrace.Client
}

var (
	numberOfRetryAttempts = 12
	retryInterval         = 5 //in seconds
)

// NewServer creates the server
func NewServer() (*Server, error) {

	var database db.Connection

	err := retry(numberOfRetryAttempts, time.Duration(retryInterval)*time.Second, func() (err error) {
		database, err = db.Create()
		return
	})

	if err != nil {
		fmt.Println("Unable to connect to database on ", err.Error())
		return nil, errors.New(err.Error())
	}

	onspotClient := onspot.New()
	socketClient := socket.New()
	appNexusClient := appnexus.New()
	paytraceClient := paytrace.New()
	if err != nil {
		return nil, err
	}

	return &Server{
		Database: &database,
		Onspot:   &onspotClient,
		Socket:   &socketClient,
		AppNexus: &appNexusClient,
		PayTrace: paytraceClient,
	}, nil
}

func (s *Server) Close() {
	s.Database.Close()
}
