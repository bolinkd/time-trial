package domain

import (
	"encoding/json"
	"errors"
	"github.com/bolinkd/time-trial/models"
)

var (
	ErrInvalidRentalOutTime AppError = errors.New("invalid rental out time")
	ErrInvalidRentalShellID AppError = errors.New("invalid shell id")
)

type Rental struct {
	*models.Rental
}

type RentalSlice struct {
	models.RentalSlice
}

func (r Rental) Validate() error {
	if !r.OutTime.Valid {
		return ErrInvalidRentalOutTime
	}
	if !r.ShellID.Valid {
		return ErrInvalidRentalShellID
	}
	return nil
}

func (r *Rental) MarshalJSON() ([]byte, error) {
	if r.R == nil {
		return json.Marshal(r.Rental)
	}
	return json.Marshal(&struct {
		*models.Rental
		Shell       *models.Shell           `json:"shell,omitempty"`
		RentalUsers models.RentalRowerSlice `json:"rental_rowers,omitempty"`
	}{
		Rental:      r.Rental,
		Shell:       r.R.Shell,
		RentalUsers: r.R.RentalRowers,
	})
}

func (rs RentalSlice) MarshalJSON() ([]byte, error) {
	rsd := make([]*Rental, 0)
	for _, r := range rs.RentalSlice {
		rsd = append(rsd, &Rental{r})
	}
	return json.Marshal(rsd)
}
