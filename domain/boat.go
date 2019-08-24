package domain

import (
	"encoding/json"
	"errors"
	"github.com/bolinkd/time-trial/models"
)

var (
	ErrInvalidBoatTimeTrial AppError = errors.New("invalid boat time trial id")
	ErrInvalidBoatMarker    AppError = errors.New("invalid boat bow marker")
)

type Boat struct {
	*models.Boat
}

type BoatSlice struct {
	models.BoatSlice
}

func (b Boat) Validate() error {
	if !b.TimeTrialID.Valid {
		return ErrInvalidBoatTimeTrial
	}
	if !b.BowMarker.Valid {
		return ErrInvalidBoatMarker
	}
	return nil
}

func (b *Boat) MarshalJSON() ([]byte, error) {
	if b.R == nil {
		return json.Marshal(b.Boat)
	}
	return json.Marshal(&struct {
		*models.Boat
		TimeTrial *models.TimeTrial `json:"time_trial,omitempty"`
	}{
		Boat:      b.Boat,
		TimeTrial: b.R.TimeTrial,
	})
}

func (bs BoatSlice) MarshalJSON() ([]byte, error) {
	var bsd []*Boat
	for _, b := range bs.BoatSlice {
		bsd = append(bsd, &Boat{b})
	}
	return json.Marshal(bsd)
}
