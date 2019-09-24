package domain

import (
	"encoding/json"
	"errors"
	"github.com/bolinkd/time-trial/models"
)

var (
	ErrInvalidRowerRentalRowerID  AppError = errors.New("invalid rower id")
	ErrInvalidRowerRentalRentalID AppError = errors.New("invalid rental id")
)

type RentalRower struct {
	*models.RentalRower
}

type RentalRowerSlice struct {
	models.RentalRowerSlice
}

func (r RentalRower) Validate() error {
	if !r.RowerID.Valid {
		return ErrInvalidRowerRentalRowerID
	}
	if !r.RentalID.Valid {
		return ErrInvalidRowerRentalRentalID
	}
	return nil
}

func (r *RentalRower) MarshalJSON() ([]byte, error) {
	if r.R == nil {
		return json.Marshal(r.RentalRower)
	}
	return json.Marshal(&struct {
		*models.RentalRower
		Rower  *models.Rower  `json:"rower,omitempty"`
		Rental *models.Rental `json:"rental,omitempty"`
	}{
		RentalRower: r.RentalRower,
		Rower:       r.R.Rower,
		Rental:      r.R.Rental,
	})
}

func (rs RentalRowerSlice) MarshalJSON() ([]byte, error) {
	rsd := make([]*RentalRower, 0)
	for _, r := range rs.RentalRowerSlice {
		rsd = append(rsd, &RentalRower{r})
	}
	return json.Marshal(rsd)
}
