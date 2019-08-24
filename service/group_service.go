package service

import (
	"database/sql"
	"github.com/bolinkd/time-trial/db"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
)

type GroupServiceInterface interface {
	GetGroupsByClub(db db.DatabaseInterface, clubID int) (models.GroupSlice, error)
	GetGroupByID(db db.DatabaseInterface, id int) (*models.Group, error)
	CreateGroup(db db.DatabaseInterface, group *models.Group) error
	UpdateGroup(db db.DatabaseInterface, group *models.Group) error
}

func (Services) GetGroupsByClub(db db.DatabaseInterface, clubID int) (models.GroupSlice, error) {
	return db.FindGroupsByClubID(nil, clubID)
}

func (Services) GetGroupByID(db db.DatabaseInterface, id int) (*models.Group, error) {
	club, err := db.FindGroupByID(nil, id)
	if err == sql.ErrNoRows {
		return nil, domain.ErrClubNotFound
	} else if err != nil {
		return nil, err
	}
	return club, err
}

func (Services) CreateGroup(db db.DatabaseInterface, group *models.Group) error {
	return db.AddGroup(nil, group)
}

func (Services) UpdateGroup(db db.DatabaseInterface, group *models.Group) error {
	return db.UpdateGroup(nil, group)
}
