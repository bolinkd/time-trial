package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrganization_Validate(t *testing.T) {
	org := Organization{}

	t.Run("organization validate - valid", func(t *testing.T) {
		err := org.Validate()
		assert.NoError(t, err)
	})
}
