package db

import (
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type ClubDBInterface interface {
	FindClubs(tx boil.Executor) (models.ClubSlice, error)
	FindClubsByOrganization(tx boil.Executor, orgID int) (models.ClubSlice, error)
	FindClubByID(tx boil.Executor, id int) (*models.Club, error)
	AddClub(tx boil.Executor, club *models.Club) error
	UpdateClub(tx boil.Executor, club *models.Club) error
}

func (conn Connection) FindClubs(tx boil.Executor) (models.ClubSlice, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.Clubs().All(tx)
}

func (conn Connection) FindClubsByOrganization(tx boil.Executor, orgID int) (models.ClubSlice, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.Clubs(qm.Where("organization_id = ?", orgID)).All(tx)
}

func (conn Connection) FindClubByID(tx boil.Executor, id int) (*models.Club, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.Clubs(qm.Where("id = ?", id)).One(tx)
}

func (conn Connection) AddClub(tx boil.Executor, club *models.Club) error {
	if tx == nil {
		tx = conn.DB
	}
	return club.Insert(tx, boil.Infer())
}

func (conn Connection) UpdateClub(tx boil.Executor, club *models.Club) error {
	if tx == nil {
		tx = conn.DB
	}

	rowsAff, err := club.Update(tx, boil.Infer())
	if err != nil {
		return err
	}

	if rowsAff == 0 {
		return domain.ErrClubNotFound
	}

	return nil
}
