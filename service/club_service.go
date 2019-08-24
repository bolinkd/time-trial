package service

import (
	"database/sql"
	"github.com/bolinkd/time-trial/db"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
)

type ClubServiceInterface interface {
	GetClubsByOrganization(db db.DatabaseInterface, organizationID int) (models.ClubSlice, error)
	GetClubByID(db db.DatabaseInterface, clubID int) (*models.Club, error)
	CreateClub(db db.DatabaseInterface, club *models.Club) error
	UpdateClub(db db.DatabaseInterface, club *models.Club) error
}

func (Services) GetClubsByOrganization(db db.DatabaseInterface, organizationID int) (models.ClubSlice, error) {
	return db.FindClubsByOrganization(nil, organizationID)
}

func (Services) GetClubByID(db db.DatabaseInterface, clubID int) (*models.Club, error) {
	club, err := db.FindClubByID(nil, clubID)
	if err == sql.ErrNoRows {
		return nil, domain.ErrClubNotFound
	} else if err != nil {
		return nil, err
	}
	return club, err
}

func (Services) CreateClub(db db.DatabaseInterface, club *models.Club) error {
	return db.AddClub(nil, club)
}

func (Services) UpdateClub(db db.DatabaseInterface, club *models.Club) error {
	return db.UpdateClub(nil, club)
}
