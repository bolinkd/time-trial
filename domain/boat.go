package domain

import (
	"errors"
	"gopkg.in/volatiletech/null.v6"
	"time"
)

var (
	ErrInvalidBoat = errors.New("invalid boat")
)

type Boat struct {
	ID          int         `json:"id"`
	TimeTrialID null.Int    `json:"time_trial_id"`
	Name        null.String `json:"name"`
	Start       null.Int    `json:"start"`
	End         null.Int    `json:"end"`
	Time        null.Int    `json:"time"`
	BowMarker   null.Int    `json:"bow_marker"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type BoatSlice []Boat

func (b Boat) Validate() error {
	if !b.TimeTrialID.Valid {
		return ErrInvalidBoat
	}
	if !b.BowMarker.Valid {
		return ErrInvalidBoat
	}
	return nil
}
