// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import boil "github.com/volatiletech/sqlboiler/boil"

import mock "github.com/stretchr/testify/mock"
import models "github.com/businessinstincts/traxone/models"

// UserDBInterface is an autogenerated mock type for the UserDBInterface type
type UserDBInterface struct {
	mock.Mock
}

// AddUser provides a mock function with given fields: traxUser, tx
func (_m *UserDBInterface) AddUser(traxUser *models.Traxuser, tx boil.Executor) error {
	ret := _m.Called(traxUser, tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Traxuser, boil.Executor) error); ok {
		r0 = rf(traxUser, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAllUsersByOrgID provides a mock function with given fields: orgID, tx
func (_m *UserDBInterface) FindAllUsersByOrgID(orgID int, tx boil.Executor) (models.TraxuserSlice, error) {
	ret := _m.Called(orgID, tx)

	var r0 models.TraxuserSlice
	if rf, ok := ret.Get(0).(func(int, boil.Executor) models.TraxuserSlice); ok {
		r0 = rf(orgID, tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.TraxuserSlice)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, boil.Executor) error); ok {
		r1 = rf(orgID, tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByID provides a mock function with given fields: userID, tx
func (_m *UserDBInterface) FindUserByID(userID int, tx boil.Executor) (*models.Traxuser, error) {
	ret := _m.Called(userID, tx)

	var r0 *models.Traxuser
	if rf, ok := ret.Get(0).(func(int, boil.Executor) *models.Traxuser); ok {
		r0 = rf(userID, tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Traxuser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, boil.Executor) error); ok {
		r1 = rf(userID, tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: traxUser, tx
func (_m *UserDBInterface) UpdateUser(traxUser *models.Traxuser, tx boil.Executor) error {
	ret := _m.Called(traxUser, tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Traxuser, boil.Executor) error); ok {
		r0 = rf(traxUser, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
