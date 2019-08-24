package db

import (
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"time"
)

type RentalDBInterface interface {
	FindRentals(tx boil.Executor, orgID int, isActive bool, startTime time.Time, endTime time.Time) (models.RentalSlice, error)
	FindRentalsByShellID(tx boil.Executor, boatID int) (models.RentalSlice, error)
	FindRentalByID(tx boil.Executor, id int) (*models.Rental, error)
	AddRental(tx boil.Executor, boat *models.Rental) error
	UpdateRental(tx boil.Executor, boat *models.Rental) error
	DeleteRental(tx boil.Executor, id int) error
}

func (conn Connection) FindRentals(tx boil.Executor, orgID int, isActive bool, startTime time.Time, endTime time.Time) (models.RentalSlice, error) {
	if tx == nil {
		tx = conn.DB
	}

	queries := []qm.QueryMod{
		qm.InnerJoin("shell s on rental.shell_id = s.id"),
		qm.InnerJoin("club c on s.club_id = c.id"),
		qm.InnerJoin("organization o on c.organization_id = o.id"),
		qm.Where("o.id = ?", orgID),
	}

	if isActive {
		queries = append(queries, models.RentalWhere.InTime.IsNull())
	}
	queries = append(queries, models.RentalWhere.OutTime.GTE(null.TimeFrom(startTime)))
	queries = append(queries, models.RentalWhere.OutTime.LTE(null.TimeFrom(endTime)))

	return models.Rentals(queries...).All(tx)
}

func (conn Connection) FindRentalsByShellID(tx boil.Executor, shellID int) (models.RentalSlice, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.Rentals(qm.Where("shell_id = ?", shellID)).All(tx)
}

func (conn Connection) FindRentalByID(tx boil.Executor, id int) (*models.Rental, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.Rentals(qm.Where("id = ?", id)).One(tx)
}

func (conn Connection) AddRental(tx boil.Executor, rental *models.Rental) error {
	if tx == nil {
		tx = conn.DB
	}
	return rental.Insert(tx, boil.Infer())
}

func (conn Connection) UpdateRental(tx boil.Executor, rental *models.Rental) error {
	if tx == nil {
		tx = conn.DB
	}

	rowsAff, err := rental.Update(tx, boil.Infer())
	if err != nil {
		return err
	}

	if rowsAff == 0 {
		return domain.ErrRentalNotFound
	}

	return nil
}

func (conn Connection) DeleteRental(tx boil.Executor, id int) error {
	if tx == nil {
		tx = conn.DB
	}
	rowsAff, err := models.Rentals(qm.Where("id = ?", id)).DeleteAll(tx)
	if err != nil {
		return err
	}

	if rowsAff == 0 {
		return domain.ErrRentalNotFound
	}

	return nil
}
