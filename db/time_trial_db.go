package db

import (
	"github.com/businessinstincts/traxone/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type UserDBInterface interface {
	FindUserByID(userID int, tx boil.Executor) (*models.Traxuser, error)
	AddUser(traxUser *models.Traxuser, tx boil.Executor) error
	FindAllUsersByOrgID(orgID int, tx boil.Executor) (models.TraxuserSlice, error)
	UpdateUser(traxUser *models.Traxuser, tx boil.Executor) error
}

func (conn Connection) FindUserByID(userID int, tx boil.Executor) (*models.Traxuser, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.Traxusers(tx, qm.Where("id = ?", userID), qm.Load("Org")).One()
}

func (conn Connection) AddUser(traxUser *models.Traxuser, tx boil.Executor) error {
	if tx == nil {
		tx = conn.DB
	}
	return traxUser.Insert(tx)
}

func (conn Connection) UpdateUser(traxUser *models.Traxuser, tx boil.Executor) error {
	if tx == nil {
		tx = conn.DB
	}
	return traxUser.Update(tx)
}

func (conn Connection) FindAllUsersByOrgID(orgID int, tx boil.Executor) (models.TraxuserSlice, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.Traxusers(tx, qm.Where("org_id = ?", orgID), qm.Load("Org")).All()
}
