package service

import (
	"database/sql"
	"github.com/bolinkd/time-trial/db"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type RentalRowerServiceInterface interface {
	GetRentalRowersByRental(db db.DatabaseInterface, id int) (models.RentalRowerSlice, error)
	CreateRentalRower(db db.DatabaseInterface, rentalRower *models.RentalRower) error
	UpdateRentalRower(db db.DatabaseInterface, rentalRower *models.RentalRower) error
	DeleteRentalRower(db db.DatabaseInterface, id int) error
}

func (Services) GetRentalRowersByRental(db db.DatabaseInterface, id int) (models.RentalRowerSlice, error) {
	loadRels := []qm.QueryMod{
		qm.Load(models.RentalRowerRels.Rower),
		qm.Load(qm.Rels(models.RentalRowerRels.Rental, models.RentalRels.Shell)),
	}
	rentalRowers, err := db.FindRentalRowersByRentalID(nil, id, loadRels)
	if err == sql.ErrNoRows {
		return nil, domain.ErrRentalNotFound
	} else if err != nil {
		return nil, err
	}

	return rentalRowers, err
}

func (Services) CreateRentalRower(db db.DatabaseInterface, rentalRower *models.RentalRower) error {
	return db.AddRentalRower(nil, rentalRower)
}

func (Services) UpdateRentalRower(db db.DatabaseInterface, rentalRower *models.RentalRower) error {
	return db.UpdateRentalRower(nil, rentalRower)
}

func (Services) DeleteRentalRower(db db.DatabaseInterface, id int) error {
	return db.DeleteRentalRower(nil, id)
}
