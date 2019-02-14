package service

import (
	"errors"
	"github.com/bolinkd/time-trial/domain"
)

var (
	ErrTimeTrialNotFound domain.AppError = errors.New("time trial not found")
	ErrBoatNotFound      domain.AppError = errors.New("boat not found")
)
