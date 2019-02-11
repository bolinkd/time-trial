package modext

import (
	"github.com/businessinstincts/traxone/domain"
	"github.com/businessinstincts/traxone/models"
)

func ConvertOrganizationsToDomain(orgs models.OrganizationSlice) (domain.OrganizationSlice, error) {
	var os = domain.OrganizationSlice{}
	for _, org := range orgs {
		o := ConvertOrganizationToDomain(org)
		os = append(os, *o)

	}
	return os, nil
}

func ConvertOrganizationToDomain(org *models.Organization) *domain.Organization {
	return &domain.Organization{
		ID:          org.ID,
		Description: org.Description,
		CreatedAt:   org.CreatedAt,
		UpdatedAt:   org.UpdatedAt,
	}
}

func ConvertOrganizationToModel(org domain.Organization) *models.Organization {
	return &models.Organization{
		ID:          org.ID,
		Description: org.Description,
		CreatedAt:   org.CreatedAt,
		UpdatedAt:   org.UpdatedAt,
	}
}
