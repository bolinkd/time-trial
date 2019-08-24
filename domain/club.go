package domain

import (
	"encoding/json"
	"errors"
	"github.com/bolinkd/time-trial/models"
)

var (
	ErrInvalidClubOrganizationId AppError = errors.New("invalid club organization id")
	ErrInvalidClubName           AppError = errors.New("invalid club name")
	ErrInvalidClubAbbreviation   AppError = errors.New("invalid club abbreviation")
)

type Club struct {
	*models.Club
}

type ClubSlice struct {
	models.ClubSlice
}

func (c Club) Validate() error {
	if !c.OrganizationID.Valid {
		return ErrInvalidClubOrganizationId
	}
	if !c.Name.Valid {
		return ErrInvalidClubName
	}
	if !c.Abbreviation.Valid {
		return ErrInvalidClubAbbreviation
	}
	return nil
}

func (c *Club) MarshalJSON() ([]byte, error) {
	if c.R == nil {
		return json.Marshal(c.Club)
	}
	return json.Marshal(&struct {
		*models.Club
		Shells       models.ShellSlice    `json:"shells,omitempty"`
		Groups       models.GroupSlice    `json:"groups,omitempty"`
		Organization *models.Organization `json:"organization,omitempty"`
	}{
		Club:         c.Club,
		Shells:       c.R.Shells,
		Groups:       c.R.Groups,
		Organization: c.R.Organization,
	})
}

func (cs ClubSlice) MarshalJSON() ([]byte, error) {
	var csd []*Club
	for _, c := range cs.ClubSlice {
		csd = append(csd, &Club{c})
	}
	return json.Marshal(csd)
}
