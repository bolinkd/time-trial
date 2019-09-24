package domain

import (
	"encoding/json"
	"errors"
	"github.com/bolinkd/time-trial/models"
)

var (
	ErrInvalidShellName           AppError = errors.New("invalid shell name")
	ErrInvalidShellGroupID        AppError = errors.New("invalid shell group id")
	ErrInvalidShellType           AppError = errors.New("invalid shell type")
	ErrInvalidShellTypeOutOfRange AppError = errors.New("invalid shell type - not in range")
)

type ShellType int

const (
	ShellTypeSingle ShellType = iota
	ShellTypeDouble
	ShellTypeQuad
	ShellTypeEight
)

type Shell struct {
	*models.Shell
}

type ShellSlice struct {
	models.ShellSlice
}

func (s Shell) Validate() error {
	if !s.Name.Valid {
		return ErrInvalidShellName
	}
	if !s.GroupID.Valid {
		return ErrInvalidShellGroupID
	}
	if !s.Type.Valid {
		return ErrInvalidShellType
	} else if s.Type.Int < int(ShellTypeSingle) || s.Type.Int > int(ShellTypeEight) {
		return ErrInvalidShellTypeOutOfRange
	}
	return nil
}

func (s *Shell) MarshalJSON() ([]byte, error) {
	if s.R == nil {
		return json.Marshal(s.Shell)
	}
	return json.Marshal(&struct {
		*models.Shell
		Group   *models.Group      `json:"group,omitempty"`
		Rentals models.RentalSlice `json:"rentals,omitempty"`
	}{
		Shell:   s.Shell,
		Group:   s.R.Group,
		Rentals: s.R.Rentals,
	})
}

func (ss ShellSlice) MarshalJSON() ([]byte, error) {
	ssd := make([]*Shell, 0)
	for _, s := range ss.ShellSlice {
		ssd = append(ssd, &Shell{s})
	}
	return json.Marshal(ssd)
}
