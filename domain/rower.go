package domain

import (
	"encoding/json"
	"errors"
	"github.com/bolinkd/time-trial/models"
)

var (
	ErrInvalidRowerFirstName AppError = errors.New("invalid rower first name")
	ErrInvalidRowerLastName  AppError = errors.New("invalid rower last name")
	ErrInvalidRowerGroupID   AppError = errors.New("invalid rower group id")
)

type Rower struct {
	*models.Rower
}

type RowerSlice struct {
	models.RowerSlice
}

func (r Rower) Validate() error {
	if !r.FirstName.Valid {
		return ErrInvalidRowerFirstName
	}
	if !r.LastName.Valid {
		return ErrInvalidRowerLastName
	}
	if !r.GroupID.Valid {
		return ErrInvalidRowerGroupID
	}
	return nil
}

func (r *Rower) MarshalJSON() ([]byte, error) {
	if r.R == nil {
		return json.Marshal(r.Rower)
	}
	return json.Marshal(&struct {
		*models.Rower
		RentalRowers models.RentalRowerSlice `json:"rental_rowers,omitempty"`
		Group        *models.Group           `json:"group,omitempty"`
	}{
		Rower:        r.Rower,
		RentalRowers: r.R.RentalRowers,
		Group:        r.R.Group,
	})
}

func (rs RowerSlice) MarshalJSON() ([]byte, error) {
	var rsd []*Rower
	for _, r := range rs.RowerSlice {
		rsd = append(rsd, &Rower{r})
	}
	return json.Marshal(rsd)
}
