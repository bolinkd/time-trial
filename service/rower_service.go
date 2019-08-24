package service

import (
	"database/sql"
	"github.com/bolinkd/time-trial/db"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
)

type RowerServiceInterface interface {
	GetRowersByGroupID(db db.DatabaseInterface, groupID int) (models.RowerSlice, error)
	GetRowerByID(db db.DatabaseInterface, id int) (*models.Rower, error)
	CreateRower(db db.DatabaseInterface, rower *models.Rower) error
	UpdateRower(db db.DatabaseInterface, rower *models.Rower) error
}

func (Services) GetRowersByGroupID(db db.DatabaseInterface, groupID int) (models.RowerSlice, error) {
	return db.FindRowersByGroupID(nil, groupID)
}

func (Services) GetRowerByID(db db.DatabaseInterface, id int) (*models.Rower, error) {
	rower, err := db.FindRowerByID(nil, id)
	if err == sql.ErrNoRows {
		return nil, domain.ErrRowerNotFound
	} else if err != nil {
		return nil, err
	}
	return rower, nil
}

func (Services) CreateRower(db db.DatabaseInterface, rower *models.Rower) error {
	return db.AddRower(nil, rower)
}

func (Services) UpdateRower(db db.DatabaseInterface, rower *models.Rower) error {
	return db.UpdateRower(nil, rower)
}
