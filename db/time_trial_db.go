package db

import (
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type TimeTrialDBInterface interface {
	FindTimeTrials(tx boil.Executor) (models.TimeTrialSlice, error)
	FindTimeTrialByID(id int, tx boil.Executor) (*models.TimeTrial, error)
	AddTimeTrial(timeTrial *models.TimeTrial, tx boil.Executor) error
	UpdateTimeTrial(timeTrial *models.TimeTrial, tx boil.Executor) error
}

func (conn Connection) FindTimeTrials(tx boil.Executor) (models.TimeTrialSlice, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.TimeTrials(qm.Load(models.TimeTrialRels.Boats)).All(tx)
}

func (conn Connection) FindTimeTrialByID(id int, tx boil.Executor) (*models.TimeTrial, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.TimeTrials(qm.Where("id = ?", id), qm.Load(models.TimeTrialRels.Boats)).One(tx)
}

func (conn Connection) AddTimeTrial(timeTrial *models.TimeTrial, tx boil.Executor) error {
	if tx == nil {
		tx = conn.DB
	}
	return timeTrial.Insert(tx, boil.Infer())
}

func (conn Connection) UpdateTimeTrial(timeTrial *models.TimeTrial, tx boil.Executor) error {
	if tx == nil {
		tx = conn.DB
	}
	rowsAff, err := timeTrial.Update(tx, boil.Infer())
	if err != nil {
		return err
	}

	if rowsAff != 0 {
		return domain.ErrTimeTrialNotFound
	}
	return err
}
