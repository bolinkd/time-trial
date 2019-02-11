package db

import (
	"errors"
	"fmt"
	"github.com/volatiletech/sqlboiler/boil"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/level48/cafe-service/setup"
)

const (
	defaultHost = "localhost"
	defaultPort = "5432"
	defaultDB   = "time_trial_db"
	defaultUser = "postgres"
	defaultPW   = "nopassword"
)

// Connection contains structure with Connection connection and context
type Connection struct {
	DB *sqlx.DB
}

type DatabaseInterface interface {
	TimeTrialDBInterface
	GetTransactor() (boil.Transactor, error)
}

var (
	host     = setup.EnvsString([]string{"RDS_HOSTNAME", "LOCATION_POSTGRES_1_PORT_5432_TCP_ADDR"}, defaultHost)
	port     = setup.EnvsString([]string{"RDS_PORT", "LOCATION_POSTGRES_1_PORT_5432_TCP_PORT"}, defaultPort)
	name     = setup.EnvsString([]string{"RDS_DB_NAME"}, defaultDB)
	user     = setup.EnvsString([]string{"RDS_USERNAME"}, defaultUser)
	password = setup.EnvsString([]string{"RDS_PASSWORD", "DB_ENV_POSTGRES_PASSWORD"}, defaultPW)

	cnx = fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, name, user, password)
)

// Create creates the new db server connection
func Create() (Connection, error) {
	db, err := sqlx.Connect("postgres", cnx)
	if err != nil {
		fmt.Println("Unable to connect to database on", host, err.Error())
		return Connection{}, errors.New(err.Error())
	}

	log.Println("Connected to database...")

	// set to reasonable values for production
	db.SetMaxIdleConns(4)
	db.SetMaxOpenConns(16)

	return Connection{DB: db}, nil
}

// Close the database connection
func (conn *Connection) Close() {
	err := conn.DB.Close()
	if err != nil {
		log.Println(err)
	}
}

func (Connection) GetTransactor() (boil.Transactor, error) {
	return boil.Begin()
}
