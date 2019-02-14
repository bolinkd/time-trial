package db

import (
	"github.com/bolinkd/time-trial/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type BoatDBInterface interface {
	FindBoatsByTimeTrial(timeTrialID int, tx boil.Executor) (models.BoatSlice, error)
	FindBoatByID(id int, tx boil.Executor) (*models.Boat, error)
	AddBoat(boat *models.Boat, tx boil.Executor) error
	UpdateBoat(boat *models.Boat, tx boil.Executor) error
}

func (conn Connection) FindBoatsByTimeTrial(timeTrialID int, tx boil.Executor) (models.BoatSlice, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.Boats(tx, qm.Where("time_trial_id = ?", timeTrialID)).All()
}

func (conn Connection) FindBoatByID(id int, tx boil.Executor) (*models.Boat, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.Boats(tx, qm.Where("id = ?", id)).One()
}

func (conn Connection) AddBoat(boat *models.Boat, tx boil.Executor) error {
	if tx == nil {
		tx = conn.DB
	}
	return boat.Insert(tx)
}

func (conn Connection) UpdateBoat(boat *models.Boat, tx boil.Executor) error {
	if tx == nil {
		tx = conn.DB
	}
	return boat.Update(tx)
}
