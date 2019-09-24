package db

import (
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type RowerDBInterface interface {
	FindRowers(tx boil.Executor) (models.RowerSlice, error)
	FindRowersByOrganizationID(tx boil.Executor, orgID int) (models.RowerSlice, error)
	FindRowersByGroupID(tx boil.Executor, groupID int) (models.RowerSlice, error)
	FindRowerByID(tx boil.Executor, id int) (*models.Rower, error)
	AddRower(tx boil.Executor, rower *models.Rower) error
	UpdateRower(tx boil.Executor, rower *models.Rower) error
}

func (conn Connection) FindRowers(tx boil.Executor) (models.RowerSlice, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.Rowers().All(tx)
}

func (conn Connection) FindRowersByOrganizationID(tx boil.Executor, orgID int) (models.RowerSlice, error) {
	if tx == nil {
		tx = conn.DB
	}

	return models.Rowers(models.RowerWhere.OrganizationID.EQ(null.IntFrom(orgID))).All(tx)
}

func (conn Connection) FindRowersByGroupID(tx boil.Executor, groupID int) (models.RowerSlice, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.Rowers(qm.Where("group_id = ?", groupID)).All(tx)
}

func (conn Connection) FindRowerByID(tx boil.Executor, id int) (*models.Rower, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.Rowers(qm.Where("id = ?", id)).One(tx)
}

func (conn Connection) AddRower(tx boil.Executor, rower *models.Rower) error {
	if tx == nil {
		tx = conn.DB
	}
	return rower.Insert(tx, boil.Infer())
}

func (conn Connection) UpdateRower(tx boil.Executor, rower *models.Rower) error {
	if tx == nil {
		tx = conn.DB
	}
	rowsAff, err := rower.Update(tx, boil.Infer())
	if err != nil {
		return err
	}

	if rowsAff == 0 {
		return domain.ErrRowerNotFound
	}
	return err
}
