package domain

import (
	"errors"
	"gopkg.in/volatiletech/null.v6"
	"time"
)

var (
	ErrInvalidTraxUserEmail = errors.New("invalid user email")
)

const (
	RoleStandard Role = iota
	RoleAdmin
)

type Role int

type UserCreation struct {
	User         Traxuser     `json:"user"`
	Organization Organization `json:"organization"`
}

type Traxuser struct {
	ID        int         `json:"id"`
	FirstName null.String `json:"first_name"`
	LastName  null.String `json:"last_name"`
	Email     string      `json:"email"`
	OrgID     int         `json:"org_id"`
	Role      Role        `json:"role"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type TraxuserSlice []Traxuser

func (u Traxuser) Validate() error {
	if u.Email == "" {
		return ErrInvalidTraxUserEmail
	}
	return nil
}

func (u UserCreation) Validate() error {
	err := u.User.Validate()
	if err != nil {
		return err
	}
	err = u.Organization.Validate()
	if err != nil {
		return err
	}
	return nil
}
