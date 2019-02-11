package service

import (
	"errors"
	"github.com/businessinstincts/traxone/domain"
)

var (
	ErrTimeTrialNotFound domain.TraxError = errors.New("timetrial not found")
)
