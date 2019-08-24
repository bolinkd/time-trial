package db

import (
	"database/sql"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type AuthDBInterface interface {
	ValidateAuthToken(tx boil.Executor, orgID int, token string) (bool, error)
	GetToken(tx boil.Executor, auth *models.OrganizationAuth) (string, error)
}

func (conn Connection) ValidateAuthToken(tx boil.Executor, orgID int, token string) (bool, error) {
	if tx == nil {
		tx = conn.DB
	}
	return models.OrganizationAuths(qm.Where("organization_id = ? and token = ?", orgID, token)).Exists(tx)
}

func (conn Connection) GetToken(tx boil.Executor, auth *models.OrganizationAuth) (string, error) {
	if tx == nil {
		tx = conn.DB
	}
	orgAuth, err := models.OrganizationAuths(qm.Where("organization_id = ? and phrase = ?", auth.OrganizationID, auth.Phrase)).One(tx)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", domain.ErrAuthNotFound
		}
		return "", err
	}

	return orgAuth.Token.String, nil
}
