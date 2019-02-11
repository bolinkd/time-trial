package modext

import (
	"github.com/businessinstincts/traxone/domain"
	"github.com/businessinstincts/traxone/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertUsersToDomain(t *testing.T) {
	var usersM models.TraxuserSlice

	t.Run("empty/nil array", func(t *testing.T) {
		users := ConvertUsersToDomain(usersM)
		assert.Equal(t, users, domain.TraxuserSlice{})
	})
	usersM = append(usersM, &models.Traxuser{})
	t.Run("non-empty", func(t *testing.T) {
		users := ConvertUsersToDomain(usersM)
		assert.NotNil(t, users)
	})
}

func TestConvertUserToDomain(t *testing.T) {
	userM := &models.Traxuser{}
	userD := &domain.Traxuser{}

	t.Run("empty user", func(t *testing.T) {
		actual := ConvertUserToDomain(userM)
		assert.Equal(t, actual, userD)
	})
}

func TestConvertUserToModel(t *testing.T) {
	userM := &models.Traxuser{}
	userD := domain.Traxuser{}

	actual := ConvertUserToModel(userD)

	assert.Equal(t, actual, userM)
}
