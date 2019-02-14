package service

import (
	"errors"
	"github.com/businessinstincts/traxone/domain"
)

var (
	ErrTimeTrialNotFound domain.TraxError = errors.New("time trial not found")
	ErrBoatNotFound      domain.TraxError = errors.New("boat not found")
)
