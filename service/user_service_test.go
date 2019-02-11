package service

import (
	"github.com/businessinstincts/traxone/domain"
	"github.com/businessinstincts/traxone/utils/mocks"
	"github.com/stretchr/testify/assert"
	"gopkg.in/volatiletech/null.v6"
	"testing"
)

var (
	userService = UserService{}
)

func TestUserService_GetUserByID(t *testing.T) {

	t.Run("User Not Found", func(t *testing.T) {
		user, err := userService.GetUserByID(DBMock, mocks.UserNotFound)
		assert.Error(t, err)
		assert.Equal(t, ErrUserNotFound, err)
		assert.Nil(t, user)
	})

	t.Run("User Unknown Error", func(t *testing.T) {
		user, err := userService.GetUserByID(DBMock, mocks.UserUnknownError)
		assert.Error(t, err)
		assert.Equal(t, mocks.ErrUnknown, err)
		assert.Nil(t, user)
	})

	t.Run("Success", func(t *testing.T) {
		user, err := userService.GetUserByID(DBMock, mocks.UserStandard1)
		assert.NoError(t, err)
		assert.NotNil(t, user)
	})

}

func TestUserService_AddUser(t *testing.T) {
	userD := domain.Traxuser{
		ID:        mocks.UserUnknownError,
		OrgID:     mocks.OrgUnknownError,
		FirstName: null.StringFrom(""),
		LastName:  null.StringFrom(""),
		Email:     "",
		Role:      0,
	}

	t.Run("Org Unknown Error", func(t *testing.T) {
		addedUser, err := userService.AddUser(DBMock, userD, nil)
		assert.Error(t, err)
		assert.Equal(t, mocks.ErrUnknown, err)
		assert.Nil(t, addedUser)
	})

	userD.OrgID = mocks.OrgNotFound

	t.Run("Org Not Found", func(t *testing.T) {
		addedUser, err := userService.AddUser(DBMock, userD, nil)
		assert.Error(t, err)
		assert.Equal(t, ErrOrgNotFound, err)
		assert.Nil(t, addedUser)
	})

	userD.OrgID = mocks.OrgStandard1

	t.Run("User Unknown Error", func(t *testing.T) {
		addedUser, err := userService.AddUser(DBMock, userD, nil)
		assert.Error(t, err)
		assert.Equal(t, mocks.ErrUnknown, err)
		assert.Nil(t, addedUser)
	})

	userD.ID = mocks.UserStandard1

	t.Run("Success", func(t *testing.T) {
		addedUser, err := userService.AddUser(DBMock, userD, nil)
		assert.NoError(t, err)
		assert.NotNil(t, addedUser)
	})
}

func TestUserService_UpdateUser(t *testing.T) {
	userD := domain.Traxuser{
		ID:        mocks.UserNotFound,
		OrgID:     mocks.OrgUnknownError,
		FirstName: null.StringFrom(""),
		LastName:  null.StringFrom(""),
		Email:     "",
		Role:      0,
	}

	t.Run("User Not Found", func(t *testing.T) {
		updatedUser, err := userService.UpdateUser(DBMock, userD)
		assert.Error(t, err)
		assert.Equal(t, ErrUserNotFound, err)
		assert.Nil(t, updatedUser)
	})

	userD.ID = mocks.UserUnknownError

	t.Run("User Unknown Error", func(t *testing.T) {
		updatedUser, err := userService.UpdateUser(DBMock, userD)
		assert.Error(t, err)
		assert.Equal(t, mocks.ErrUnknown, err)
		assert.Nil(t, updatedUser)
	})

	userD.ID = mocks.UserStandard1

	t.Run("Org Unknown Error", func(t *testing.T) {
		updatedUser, err := userService.UpdateUser(DBMock, userD)
		assert.Error(t, err)
		assert.Equal(t, mocks.ErrUnknown, err)
		assert.Nil(t, updatedUser)
	})

	userD.OrgID = mocks.OrgNotFound

	t.Run("Org Not Found", func(t *testing.T) {
		updatedUser, err := userService.UpdateUser(DBMock, userD)
		assert.Error(t, err)
		assert.Equal(t, ErrOrgNotFound, err)
		assert.Nil(t, updatedUser)
	})

	userD.OrgID = mocks.OrgStandard1
	userD.ID = mocks.UserUnknownError

	t.Run("User Update Error", func(t *testing.T) {
		updatedUser, err := userService.UpdateUser(DBMock, userD)
		assert.Error(t, err)
		assert.Equal(t, mocks.ErrUnknown, err)
		assert.Nil(t, updatedUser)
	})

	userD.ID = mocks.UserStandard1

	t.Run("Success", func(t *testing.T) {
		updatedUser, err := userService.UpdateUser(DBMock, userD)
		assert.NoError(t, err)
		assert.NotNil(t, updatedUser)
	})
}
