package modext

import (
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
	"gopkg.in/volatiletech/null.v6"
)

func ConvertTimeTrialsToDomain(timeTrials models.TimeTrialSlice) domain.TimeTrialSlice {
	var us = domain.TimeTrialSlice{}
	for _, timeTrial := range timeTrials {
		u := ConvertTimeTrialToDomain(timeTrial)
		us = append(us, *u)

	}
	return us
}

func ConvertTimeTrialToDomain(timeTrial *models.TimeTrial) *domain.TimeTrial {
	return &domain.TimeTrial{
		ID:           timeTrial.ID,
		Date:         null.TimeFrom(timeTrial.Date),
		StartTime:    timeTrial.StartTime,
		EndTime:      timeTrial.EndTime,
		TimingStatus: domain.TimingStatus(timeTrial.TimingStatus.Int),
		Timers:       timeTrial.Timers.Int,
		Distance:     timeTrial.Distance.Float64,
		CreatedAt:    timeTrial.CreatedAt.Time,
		UpdatedAt:    timeTrial.UpdatedAt,
	}
}

func ConvertTimeTrialToModel(timeTrial domain.TimeTrial) *models.TimeTrial {
	return &models.TimeTrial{
		ID:           timeTrial.ID,
		Date:         timeTrial.Date.Time,
		StartTime:    timeTrial.StartTime,
		EndTime:      timeTrial.EndTime,
		TimingStatus: null.IntFrom(int(timeTrial.TimingStatus)),
		Timers:       null.IntFrom(timeTrial.Timers),
		Distance:     null.Float64From(timeTrial.Distance),
		CreatedAt:    null.TimeFrom(timeTrial.CreatedAt),
		UpdatedAt:    timeTrial.UpdatedAt,
	}
}
