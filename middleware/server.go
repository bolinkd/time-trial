package middleware

import (
	"errors"
	"fmt"
	"github.com/bolinkd/time-trial/darksky"
	"github.com/bolinkd/time-trial/db"
	"github.com/bolinkd/time-trial/service"
	"github.com/bolinkd/time-trial/socket"
	"time"
)

// Server structure
type Server struct {
	Database *db.Connection
	Services *service.Services
	Darksky  darksky.Client
	Socket   *socket.Client
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

	darkskyClient := darksky.Create()
	socketClient := socket.New()
	services := service.Create()

	return &Server{
		Database: &database,
		Darksky:  darkskyClient,
		Socket:   &socketClient,
		Services: &services,
	}, nil
}

func (s *Server) Close() {
	s.Database.Close()
	s.Services.Close()
}
