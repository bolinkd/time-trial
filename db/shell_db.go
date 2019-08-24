package db

import (
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type ShellDBInterface interface {
	FindShellsByClubID(tx boil.Executor, clubID int) (models.ShellSlice, error)
	FindShellByID(tx boil.Executor, id int) (*models.Shell, error)
	AddShell(tx boil.Executor, shell *models.Shell) error
	UpdateShell(tx boil.Executor, shell *models.Shell) error
}

func (conn Connection) FindShellsByClubID(tx boil.Executor, clubID int) (models.ShellSlice, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.Shells(qm.Where("club_id = ?", clubID)).All(tx)
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
