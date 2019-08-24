package db

import (
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type BoatDBInterface interface {
	FindBoatsByTimeTrial(tx boil.Executor, timeTrialID int) (models.BoatSlice, error)
	FindBoatByID(tx boil.Executor, id int) (*models.Boat, error)
	AddBoat(tx boil.Executor, boat *models.Boat) error
	UpdateBoat(tx boil.Executor, boat *models.Boat) error
}

func (conn Connection) FindBoatsByTimeTrial(tx boil.Executor, timeTrialID int) (models.BoatSlice, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.Boats(qm.Where("time_trial_id = ?", timeTrialID)).All(tx)
}

func (conn Connection) FindBoatByID(tx boil.Executor, id int) (*models.Boat, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.Boats(qm.Where("id = ?", id)).One(tx)
}

func (conn Connection) AddBoat(tx boil.Executor, boat *models.Boat) error {
	if tx == nil {
		tx = conn.DB
	}
	return boat.Insert(tx, boil.Infer())
}

func (conn Connection) UpdateBoat(tx boil.Executor, boat *models.Boat) error {
	if tx == nil {
		tx = conn.DB
	}

	rowsAff, err := boat.Update(tx, boil.Infer())
	if err != nil {
		return err
	}

	if rowsAff == 0 {
		return domain.ErrBoatNotFound
	}

	return nil
}
