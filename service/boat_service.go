package service

import (
	"database/sql"
	"github.com/bolinkd/time-trial/db"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/modext"
)

type BoatServiceInterface interface {
	GetBoatsForTimeTrial(db db.DatabaseInterface, timeTrialID int) (domain.BoatSlice, error)
	GetBoatByID(db db.DatabaseInterface, boatID int) (*domain.Boat, error)
	CreateBoat(db db.DatabaseInterface, boat domain.Boat) (*domain.Boat, error)
	UpdateBoat(db db.DatabaseInterface, boat domain.Boat) (*domain.Boat, error)
}

type BoatService struct{}

func (BoatService) GetBoatsForTimeTrial(db db.DatabaseInterface, timeTrialID int) (domain.BoatSlice, error) {
	boats, err := db.FindBoatsByTimeTrial(timeTrialID, nil)
	if err != nil {
		return nil, err
	}
	return modext.ConvertBoatsToDomain(boats), nil
}

func (BoatService) GetBoatByID(db db.DatabaseInterface, boatID int) (*domain.Boat, error) {
	boat, err := db.FindBoatByID(boatID, nil)
	if err == sql.ErrNoRows {
		return nil, ErrBoatNotFound
	} else if err != nil {
		return nil, err
	}
	return modext.ConvertBoatToDomain(boat), nil
}

func (BoatService) CreateBoat(db db.DatabaseInterface, boat domain.Boat) (*domain.Boat, error) {
	boatM := modext.ConvertBoatToModel(boat)

	err := db.AddBoat(boatM, nil)
	if err != nil {
		return nil, err
	}

	return modext.ConvertBoatToDomain(boatM), nil
}

func (BoatService) UpdateBoat(db db.DatabaseInterface, boat domain.Boat) (*domain.Boat, error) {
	boatM := modext.ConvertBoatToModel(boat)

	if boatM.ID == 0 {
		return nil, ErrBoatNotFound
	}

	err := db.UpdateBoat(boatM, nil)
	if err != nil {
		return nil, err
	}
	return modext.ConvertBoatToDomain(boatM), nil
}
