package domain

import (
	"encoding/json"
	"errors"
	"github.com/bolinkd/time-trial/models"
)

var (
	ErrInvalidTimeTrialDate AppError = errors.New("invalid time trial date")
)

type TimingStatus int

const (
	StatusReset TimingStatus = iota
	StatusRunning
	StatusComplete
)

type TimeTrial struct {
	*models.TimeTrial
}

type TimeTrialSlice struct {
	models.TimeTrialSlice
}

func (tt TimeTrial) Validate() error {
	if tt.Date.IsZero() {
		return ErrInvalidTimeTrialDate
	}
	return nil
}

func (tt *TimeTrial) MarshalJSON() ([]byte, error) {
	if tt.R == nil {
		return json.Marshal(tt.TimeTrial)
	}
	return json.Marshal(&struct {
		*models.TimeTrial
		Boats models.BoatSlice `json:"boats,omitempty"`
	}{
		TimeTrial: tt.TimeTrial,
		Boats:     tt.R.Boats,
	})
}

func (tts TimeTrialSlice) MarshalJSON() ([]byte, error) {
	ttsd := make([]*TimeTrial, 0)
	for _, tt := range tts.TimeTrialSlice {
		ttsd = append(ttsd, &TimeTrial{tt})
	}
	return json.Marshal(ttsd)
}
