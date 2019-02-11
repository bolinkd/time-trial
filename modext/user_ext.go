package modext

import (
	"github.com/businessinstincts/traxone/domain"
	"github.com/businessinstincts/traxone/models"
)

func ConvertUsersToDomain(users models.TraxuserSlice) domain.TraxuserSlice {
	var us = domain.TraxuserSlice{}
	for _, user := range users {
		u := ConvertUserToDomain(user)
		us = append(us, *u)

	}
	return us
}

func ConvertUserToDomain(user *models.Traxuser) *domain.Traxuser {
	return &domain.Traxuser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		OrgID:     user.OrgID,
		Role:      domain.Role(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ConvertUserToModel(user domain.Traxuser) *models.Traxuser {
	return &models.Traxuser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      int(user.Role),
		OrgID:     user.OrgID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
