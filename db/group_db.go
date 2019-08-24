package db

import (
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type GroupDBInterface interface {
	FindGroups(tx boil.Executor) (models.GroupSlice, error)
	FindGroupsByClubID(tx boil.Executor, clubID int) (models.GroupSlice, error)
	FindGroupByID(tx boil.Executor, id int) (*models.Group, error)
	AddGroup(tx boil.Executor, group *models.Group) error
	UpdateGroup(tx boil.Executor, group *models.Group) error
}

func (conn Connection) FindGroups(tx boil.Executor) (models.GroupSlice, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.Groups().All(tx)
}

func (conn Connection) FindGroupsByClubID(tx boil.Executor, clubID int) (models.GroupSlice, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.Groups(qm.Where("club_id = ?", clubID)).All(tx)
}

func (conn Connection) FindGroupByID(tx boil.Executor, id int) (*models.Group, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.Groups(qm.Where("id = ?", id)).One(tx)
}

func (conn Connection) AddGroup(tx boil.Executor, group *models.Group) error {
	if tx == nil {
		tx = conn.DB
	}
	return group.Insert(tx, boil.Infer())
}

func (conn Connection) UpdateGroup(tx boil.Executor, group *models.Group) error {
	if tx == nil {
		tx = conn.DB
	}
	rowsAff, err := group.Update(tx, boil.Infer())
	if err != nil {
		return err
	}

	if rowsAff == 0 {
		return domain.ErrGroupNotFound
	}

	return nil
}
