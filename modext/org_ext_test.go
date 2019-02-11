package modext

import (
	"github.com/businessinstincts/traxone/domain"
	"github.com/businessinstincts/traxone/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertOrganizationsToDomain(t *testing.T) {
	orgsM := models.OrganizationSlice{
		{},
	}
	_, err := ConvertOrganizationsToDomain(orgsM)

	assert.NoError(t, err)

}

func TestConvertOrganizationToDomain(t *testing.T) {
	orgD := &domain.Organization{}
	orgM := &models.Organization{}

	actual := ConvertOrganizationToDomain(orgM)

	assert.Equal(t, orgD, actual)
}

func TestConvertOrganizationToModel(t *testing.T) {
	orgD := domain.Organization{}
	orgM := &models.Organization{}

	actual := ConvertOrganizationToModel(orgD)

	assert.Equal(t, orgM, actual)
}
