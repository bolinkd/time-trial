package domain

import (
	"errors"
	"gopkg.in/volatiletech/null.v6"
	"time"
)

var (
	ErrInvalidTimeTrial = errors.New("invalid time trial")
)

type TimingStatus int

const (
	StatusReset TimingStatus = iota
	StatusRunning
	StatusComplete
)

type TimeTrial struct {
	ID           int          `json:"id"`
	Date         null.Time    `json:"date"`
	StartTime    null.Time    `json:"start_time"`
	EndTime      null.Time    `json:"end_time"`
	TimingStatus TimingStatus `json:"timing_status"`
	Timers       int          `json:"timers"`
	Distance     float64      `json:"distance"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

type TimeTrialSlice []TimeTrial

func (tt TimeTrial) Validate() error {
	if !tt.Date.Valid {
		return ErrInvalidTimeTrial
	}
	return nil
}
