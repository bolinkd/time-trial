package service

import (
	"database/sql"
	"github.com/bolinkd/time-trial/db"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/modext"
)

type TimeTrialServiceInterface interface {
	GetTimeTrials(db db.DatabaseInterface) (domain.TimeTrialSlice, error)
	GetTimeTrialById(db db.DatabaseInterface, timeTrialID int) (*domain.TimeTrial, error)
	CreateTimeTrial(db db.DatabaseInterface, timeTrial domain.TimeTrial) (*domain.TimeTrial, error)
	UpdateTimeTrial(db db.DatabaseInterface, timeTrial domain.TimeTrial) (*domain.TimeTrial, error)
}

type TimeTrialService struct{}

func (TimeTrialService) GetTimeTrials(db db.DatabaseInterface) (domain.TimeTrialSlice, error) {
	timeTrials, err := db.FindTimeTrials(nil)
	if err != nil {
		return nil, err
	}
	return modext.ConvertTimeTrialsToDomain(timeTrials), nil
}

func (TimeTrialService) GetTimeTrialById(db db.DatabaseInterface, timeTrialID int) (*domain.TimeTrial, error) {
	user, err := db.FindTimeTrialByID(timeTrialID, nil)
	if err == sql.ErrNoRows {
		return nil, ErrTimeTrialNotFound
	} else if err != nil {
		return nil, err
	}
	return modext.ConvertTimeTrialToDomain(user), nil
}

func (TimeTrialService) CreateTimeTrial(db db.DatabaseInterface, timeTrial domain.TimeTrial) (*domain.TimeTrial, error) {
	timeTrialM := modext.ConvertTimeTrialToModel(timeTrial)

	err := db.AddTimeTrial(timeTrialM, nil)
	if err != nil {
		return nil, err
	}

	return modext.ConvertTimeTrialToDomain(timeTrialM), nil
}

func (TimeTrialService) UpdateTimeTrial(db db.DatabaseInterface, timeTrial domain.TimeTrial) (*domain.TimeTrial, error) {
	timeTrialM := modext.ConvertTimeTrialToModel(timeTrial)

	if timeTrialM.ID == 0 {
		return nil, ErrTimeTrialNotFound
	}

	err := db.UpdateTimeTrial(timeTrialM, nil)
	if err != nil {
		return nil, err
	}
	return modext.ConvertTimeTrialToDomain(timeTrialM), nil
}
