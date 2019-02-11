package service

import (
	"database/sql"
	"github.com/volatiletech/sqlboiler/boil"
)

type UserServiceInterface interface {
	GetUserByID(db db.DatabaseInterface, userID int) (*domain.Traxuser, error)
	AddUser(db db.DatabaseInterface, traxUser domain.Traxuser, tx boil.Transactor) (*domain.Traxuser, error)
	UpdateUser(db db.DatabaseInterface, traxUser domain.Traxuser) (*domain.Traxuser, error)
}

type UserService struct{}

func (UserService) GetUserByID(db db.DatabaseInterface, userID int) (*domain.Traxuser, error) {
	user, err := db.FindUserByID(userID, nil)
	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	return modext.ConvertUserToDomain(user), nil
}

func (UserService) AddUser(db db.DatabaseInterface, traxUser domain.Traxuser, tx boil.Transactor) (*domain.Traxuser, error) {
	userM := modext.ConvertUserToModel(traxUser)

	_, err := db.FindOrgByID(userM.OrgID, tx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrOrgNotFound
		} else {
			return nil, err
		}
	}

	err = db.AddUser(userM, tx)
	if err != nil {
		return nil, err
	}

	return modext.ConvertUserToDomain(userM), nil
}

func (UserService) UpdateUser(db db.DatabaseInterface, traxUser domain.Traxuser) (*domain.Traxuser, error) {
	userM := modext.ConvertUserToModel(traxUser)

	if userM.ID == 0 {
		return nil, ErrUserNotFound
	}

	_, err := db.FindOrgByID(userM.OrgID, nil)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrOrgNotFound
		} else {
			return nil, err
		}
	}

	err = db.UpdateUser(userM, nil)
	if err != nil {
		return nil, err
	}
	return modext.ConvertUserToDomain(userM), nil
}
