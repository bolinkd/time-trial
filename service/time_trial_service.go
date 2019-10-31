package service

import (
	"database/sql"

	"github.com/bolinkd/time-trial/db"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
)

type TimeTrialServiceInterface interface {
	GetTimeTrials(db db.DatabaseInterface) (models.TimeTrialSlice, error)
	GetTimeTrialById(db db.DatabaseInterface, timeTrialID int) (*models.TimeTrial, error)
	CreateTimeTrial(db db.DatabaseInterface, timeTrial *models.TimeTrial) error
	UpdateTimeTrial(db db.DatabaseInterface, timeTrial *models.TimeTrial) error
}

func (Services) GetTimeTrials(db db.DatabaseInterface) (models.TimeTrialSlice, error) {
	return db.FindTimeTrials(nil)
}

func (Services) GetTimeTrialById(db db.DatabaseInterface, timeTrialID int) (*models.TimeTrial, error) {
	timeTrial, err := db.FindTimeTrialByID(timeTrialID, nil)
	if err == sql.ErrNoRows {
		return nil, domain.ErrTimeTrialNotFound
	} else if err != nil {
		return nil, err
	}
	return timeTrial, nil
}

func (Services) CreateTimeTrial(db db.DatabaseInterface, timeTrial *models.TimeTrial) error {
	return db.AddTimeTrial(timeTrial, nil)
}

func (s Services) UpdateTimeTrial(db db.DatabaseInterface, timeTrial *models.TimeTrial) error {
	timeTrial, err := db.FindTimeTrialByID(timeTrial.ID, nil)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.ErrTimeTrialNotFound
		}
		return err
	}
	return db.UpdateTimeTrial(timeTrial, nil)
}
