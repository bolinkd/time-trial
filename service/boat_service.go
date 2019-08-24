package service

import (
	"database/sql"
	"github.com/bolinkd/time-trial/db"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
)

type BoatServiceInterface interface {
	GetBoatsForTimeTrial(db db.DatabaseInterface, timeTrialID int) (models.BoatSlice, error)
	GetBoatByID(db db.DatabaseInterface, boatID int) (*models.Boat, error)
	CreateBoat(db db.DatabaseInterface, boat *models.Boat) error
	UpdateBoat(db db.DatabaseInterface, boat *models.Boat) error
}

func (Services) GetBoatsForTimeTrial(db db.DatabaseInterface, timeTrialID int) (models.BoatSlice, error) {
	return db.FindBoatsByTimeTrial(nil, timeTrialID)
}

func (Services) GetBoatByID(db db.DatabaseInterface, boatID int) (*models.Boat, error) {
	boat, err := db.FindBoatByID(nil, boatID)
	if err == sql.ErrNoRows {
		return nil, domain.ErrBoatNotFound
	} else if err != nil {
		return nil, err
	}
	return boat, nil
}

func (Services) CreateBoat(db db.DatabaseInterface, boat *models.Boat) error {
	return db.AddBoat(nil, boat)
}

func (Services) UpdateBoat(db db.DatabaseInterface, boat *models.Boat) error {
	return db.UpdateBoat(nil, boat)
}
