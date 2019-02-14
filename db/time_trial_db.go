package db

import (
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
	return models.TimeTrials(tx, qm.Load("Boats")).All()
}

func (conn Connection) FindTimeTrialByID(id int, tx boil.Executor) (*models.TimeTrial, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.TimeTrials(tx, qm.Where("id = ?", id), qm.Load("Boats")).One()
}

func (conn Connection) AddTimeTrial(timeTrial *models.TimeTrial, tx boil.Executor) error {
	if tx == nil {
		tx = conn.DB
	}
	return timeTrial.Insert(tx)
}

func (conn Connection) UpdateTimeTrial(timeTrial *models.TimeTrial, tx boil.Executor) error {
	if tx == nil {
		tx = conn.DB
	}
	return timeTrial.Update(tx)
}
