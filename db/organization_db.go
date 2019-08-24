package db

import (
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)

type OrganizationDBInterface interface {
	FindOrganizations(tx boil.Executor) (models.OrganizationSlice, error)
	FindOrganizationByID(tx boil.Executor, id int, loadRels []QueryMod) (*models.Organization, error)
	AddOrganization(tx boil.Executor, organization *models.Organization) error
	UpdateOrganization(tx boil.Executor, organization *models.Organization) error
}

func (conn Connection) FindOrganizations(tx boil.Executor) (models.OrganizationSlice, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.Organizations().All(tx)
}

func (conn Connection) FindOrganizationByID(tx boil.Executor, id int, loadRels []QueryMod) (*models.Organization, error) {
	if tx == nil {
		tx = conn.DB
	}
	loadRels = append(loadRels, Where("id = ?", id))
	return models.Organizations(loadRels...).One(tx)
}

func (conn Connection) AddOrganization(tx boil.Executor, organization *models.Organization) error {
	if tx == nil {
		tx = conn.DB
	}
	return organization.Insert(tx, boil.Infer())
}

func (conn Connection) UpdateOrganization(tx boil.Executor, organization *models.Organization) error {
	if tx == nil {
		tx = conn.DB
	}

	rowsAff, err := organization.Update(tx, boil.Infer())
	if err != nil {
		return err
	}

	if rowsAff == 0 {
		return domain.ErrOrganizationNotFound
	}

	return err
}
