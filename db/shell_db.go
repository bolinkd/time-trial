package db

import (
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type ShellDBInterface interface {
	FindShellsByOrganizationID(tx boil.Executor, orgID int) (models.ShellSlice, error)
	FindShellsByGroupID(tx boil.Executor, groupID int) (models.ShellSlice, error)
	FindShellByID(tx boil.Executor, id int) (*models.Shell, error)
	AddShell(tx boil.Executor, shell *models.Shell) error
	UpdateShell(tx boil.Executor, shell *models.Shell) error
}

func (conn Connection) FindShellsByOrganizationID(tx boil.Executor, orgID int) (models.ShellSlice, error) {
	if tx == nil {
		tx = conn.DB
	}
	queries := []qm.QueryMod{
		qm.InnerJoin(`"Group" g on shell.group_id = g.id`),
		qm.InnerJoin("organization o on g.organization_id = o.id"),
		qm.Where("o.id = ?", orgID),
	}
	return models.Shells(queries...).All(tx)
}

func (conn Connection) FindShellsByGroupID(tx boil.Executor, groupID int) (models.ShellSlice, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.Shells(qm.Where("group_id = ?", groupID)).All(tx)
}

func (conn Connection) FindShellByID(tx boil.Executor, id int) (*models.Shell, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.Shells(qm.Where("id = ?", id)).One(tx)
}

func (conn Connection) AddShell(tx boil.Executor, shell *models.Shell) error {
	if tx == nil {
		tx = conn.DB
	}
	return shell.Insert(tx, boil.Infer())
}

func (conn Connection) UpdateShell(tx boil.Executor, shell *models.Shell) error {
	if tx == nil {
		tx = conn.DB
	}

	rowsAff, err := shell.Update(tx, boil.Infer())
	if err != nil {
		return err
	}

	if rowsAff == 0 {
		return domain.ErrShellNotFound
	}

	return nil
}
