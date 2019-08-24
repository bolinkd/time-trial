package db

import (
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type RentalRowerDBInterface interface {
	FindRentalRowersByRentalID(tx boil.Executor, rentalID int, loadRels []qm.QueryMod) (models.RentalRowerSlice, error)
	AddRentalRower(tx boil.Executor, rentalRower *models.RentalRower) error
	UpdateRentalRower(tx boil.Executor, rentalRower *models.RentalRower) error
	DeleteRentalRower(tx boil.Executor, id int) error
}

func (conn Connection) FindRentalRowersByRentalID(tx boil.Executor, rentalID int, loadRels []qm.QueryMod) (models.RentalRowerSlice, error) {
	if tx == nil {
		tx = conn.DB
	}
	loadRels = append(loadRels, qm.Where("rental_id = ?", rentalID))
	return models.RentalRowers(loadRels...).All(tx)
}

func (conn Connection) AddRentalRower(tx boil.Executor, rentalRower *models.RentalRower) error {
	if tx == nil {
		tx = conn.DB
	}
	return rentalRower.Insert(tx, boil.Infer())
}

func (conn Connection) UpdateRentalRower(tx boil.Executor, rentalRower *models.RentalRower) error {
	if tx == nil {
		tx = conn.DB
	}

	rowsAff, err := rentalRower.Update(tx, boil.Infer())
	if err != nil {
		return err
	}

	if rowsAff == 0 {
		return domain.ErrRentalRowerNotFound
	}

	return nil
}

func (conn Connection) DeleteRentalRower(tx boil.Executor, id int) error {
	if tx == nil {
		tx = conn.DB
	}
	rowsAff, err := models.RentalRowers(qm.Where("id = ?", id)).DeleteAll(tx)
	if err != nil {
		return err
	}

	if rowsAff == 0 {
		return domain.ErrRentalRowerNotFound
	}

	return nil
}
