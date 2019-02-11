package domain

import (
	"errors"
	"time"
)

var (
	ErrInvalidDescription = errors.New("invalid organization description - cannot be empty")
)

type Organization struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type OrganizationSlice []Organization

func (u Organization) Validate() error {
	return nil
}
