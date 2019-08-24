package service

import (
	"database/sql"
	"github.com/bolinkd/time-trial/db"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
)

type ShellServiceInterface interface {
	GetShellsByClub(db db.DatabaseInterface, clubID int) (models.ShellSlice, error)
	GetShellByID(db db.DatabaseInterface, shellID int) (*models.Shell, error)
	CreateShell(db db.DatabaseInterface, shell *models.Shell) error
	UpdateShell(db db.DatabaseInterface, shell *models.Shell) error
}

func (Services) GetShellsByClub(db db.DatabaseInterface, clubID int) (models.ShellSlice, error) {
	return db.FindShellsByClubID(nil, clubID)
}

func (Services) GetShellByID(db db.DatabaseInterface, shellID int) (*models.Shell, error) {
	club, err := db.FindShellByID(nil, shellID)
	if err == sql.ErrNoRows {
		return nil, domain.ErrShellNotFound
	} else if err != nil {
		return nil, err
	}
	return club, err
}

func (Services) CreateShell(db db.DatabaseInterface, shell *models.Shell) error {
	return db.AddShell(nil, shell)
}

func (Services) UpdateShell(db db.DatabaseInterface, shell *models.Shell) error {
	return db.UpdateShell(nil, shell)
}
