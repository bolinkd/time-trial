package domain

import (
	"errors"
	"github.com/bolinkd/time-trial/models"
)

var (
	ErrInvalidAuthPhrase       AppError = errors.New("invalid auth phrase")
	ErrInvalidAuthOrganization AppError = errors.New("invalid auth organization")
)

type OrganizationAuth struct {
	*models.OrganizationAuth
}

func (b OrganizationAuth) Validate() error {
	if !b.Phrase.Valid {
		return ErrInvalidAuthPhrase
	}
	if !b.OrganizationID.Valid {
		return ErrInvalidAuthOrganization
	}
	return nil
}
