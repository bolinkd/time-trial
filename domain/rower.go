package domain

import (
	"encoding/json"
	"errors"
	"github.com/bolinkd/time-trial/models"
)

var (
	ErrInvalidRowerFirstName      AppError = errors.New("invalid rower first name")
	ErrInvalidRowerLastName       AppError = errors.New("invalid rower last name")
	ErrInvalidRowerOrganizationID AppError = errors.New("invalid rower organization id")
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
	if !r.OrganizationID.Valid {
		return ErrInvalidRowerOrganizationID
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
		Organization *models.Organization    `json:"group,omitempty"`
	}{
		Rower:        r.Rower,
		RentalRowers: r.R.RentalRowers,
		Organization: r.R.Organization,
	})
}

func (rs RowerSlice) MarshalJSON() ([]byte, error) {
	rsd := make([]*Rower, 0)
	for _, r := range rs.RowerSlice {
		rsd = append(rsd, &Rower{r})
	}
	return json.Marshal(rsd)
}
