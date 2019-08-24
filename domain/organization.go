package domain

import (
	"encoding/json"
	"errors"
	"github.com/bolinkd/time-trial/models"
)

var (
	ErrInvalidOrganizationName         AppError = errors.New("invalid organization name")
	ErrInvalidOrganizationAbbreviation AppError = errors.New("invalid organization abbreviation")
)

type Organization struct {
	*models.Organization
}

type OrganizationSlice struct {
	models.OrganizationSlice
}

func (o Organization) Validate() error {
	if !o.Name.Valid {
		return ErrInvalidOrganizationName
	}
	if !o.Abbreviation.Valid {
		return ErrInvalidOrganizationAbbreviation
	}
	return nil
}

func (o *Organization) MarshalJSON() ([]byte, error) {
	if o.R == nil {
		return json.Marshal(o.Organization)
	}
	return json.Marshal(&struct {
		*models.Organization
		Clubs models.ClubSlice `json:"clubs,omitempty"`
	}{
		Organization: o.Organization,
		Clubs:        o.R.Clubs,
	})
}

func (os OrganizationSlice) MarshalJSON() ([]byte, error) {
	var osd []*Organization
	for _, o := range os.OrganizationSlice {
		osd = append(osd, &Organization{o})
	}
	return json.Marshal(osd)
}
