package service

import (
	"database/sql"
	"github.com/bolinkd/time-trial/db"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type OrganizationServiceInterface interface {
	GetOrganizations(db db.DatabaseInterface) (models.OrganizationSlice, error)
	GetOrganizationByID(db db.DatabaseInterface, organizationID int) (*models.Organization, error)
	CreateOrganization(db db.DatabaseInterface, organization *models.Organization) error
	UpdateOrganization(db db.DatabaseInterface, organization *models.Organization) error
}

func (Services) GetOrganizations(db db.DatabaseInterface) (models.OrganizationSlice, error) {
	return db.FindOrganizations(nil)
}

func (Services) GetOrganizationByID(db db.DatabaseInterface, organizationID int) (*models.Organization, error) {
	org, err := db.FindOrganizationByID(nil, organizationID, []qm.QueryMod{
		qm.Load(models.OrganizationRels.Clubs),
	})
	if err == sql.ErrNoRows {
		return nil, domain.ErrOrganizationNotFound
	} else if err != nil {
		return nil, err
	}
	return org, err
}

func (Services) CreateOrganization(db db.DatabaseInterface, organization *models.Organization) error {
	return db.AddOrganization(nil, organization)
}

func (Services) UpdateOrganization(db db.DatabaseInterface, organization *models.Organization) error {
	return db.UpdateOrganization(nil, organization)
}
