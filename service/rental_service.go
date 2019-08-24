package service

import (
	"database/sql"
	"github.com/bolinkd/time-trial/db"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
	"time"
)

type RentalServiceInterface interface {
	GetRentals(db db.DatabaseInterface, orgID int, active bool, startTime time.Time, endTime time.Time) (models.RentalSlice, error)
	GetRentalsByShell(db db.DatabaseInterface, boatID int) (models.RentalSlice, error)
	GetRentalByID(db db.DatabaseInterface, id int) (*models.Rental, error)
	CreateRental(db db.DatabaseInterface, rental *models.Rental) error
	UpdateRental(db db.DatabaseInterface, club *models.Rental) error
	DeleteRental(db db.DatabaseInterface, id int) error
}

func (Services) GetRentals(db db.DatabaseInterface, orgID int, active bool, startTime time.Time, endTime time.Time) (models.RentalSlice, error) {
	return db.FindRentals(nil, orgID, active, startTime, endTime)
}

func (Services) GetRentalsByShell(db db.DatabaseInterface, boatID int) (models.RentalSlice, error) {
	return db.FindRentalsByShellID(nil, boatID)
}

func (Services) GetRentalByID(db db.DatabaseInterface, id int) (*models.Rental, error) {
	rental, err := db.FindRentalByID(nil, id)
	if err == sql.ErrNoRows {
		return nil, domain.ErrRentalNotFound
	} else if err != nil {
		return nil, err
	}
	return rental, err
}

func (Services) CreateRental(db db.DatabaseInterface, rental *models.Rental) error {
	return db.AddRental(nil, rental)
}

func (Services) UpdateRental(db db.DatabaseInterface, rental *models.Rental) error {
	return db.UpdateRental(nil, rental)
}

func (Services) DeleteRental(db db.DatabaseInterface, id int) error {
	return db.DeleteRental(nil, id)
}
