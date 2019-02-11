package service

import (
	"errors"
	"github.com/businessinstincts/traxone/domain"
)

var (
	ErrUserNotFound     domain.TraxError = errors.New("user not found")
	ErrOrgNotFound      domain.TraxError = errors.New("organization not found")
	ErrCampaignNotFound domain.TraxError = errors.New("campaign not found")
)
