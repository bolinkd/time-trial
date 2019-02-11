package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTraxuser_Validate(t *testing.T) {
	user := Traxuser{}

	t.Run("Traxuser - missing email", func(t *testing.T) {
		err := user.Validate()

		_, isTraxErr := err.(TraxError)
		assert.True(t, isTraxErr)
		assert.Equal(t, "invalid user email", err.Error())
	})

	user.Email = "TestUser@email.ca"

	t.Run("Traxuser - valid", func(t *testing.T) {
		err := user.Validate()

		assert.NoError(t, err)
	})
}

func TestUserCreation_Validate(t *testing.T) {
	userCreation := UserCreation{}

	t.Run("Email Missing", func(t *testing.T) {
		err := userCreation.Validate()
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidTraxUserEmail, err)
	})

	userCreation.User.Email = "a@b.ca"

	t.Run("Success", func(t *testing.T) {
		err := userCreation.Validate()
		assert.NoError(t, err)
	})
}
