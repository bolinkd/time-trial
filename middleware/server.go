package middleware

import (
	"errors"
	"fmt"
	"github.com/bolinkd/time-trial/db"
	"time"
)

// Server structure
type Server struct {
	Database *db.Connection
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

	return &Server{
		Database: &database,
	}, nil
}

func (s *Server) Close() {
	s.Database.Close()
}
