package handlers

import (
	"github.com/businessinstincts/traxone/service"
)

var (
	userService     service.UserServiceInterface         = &service.UserService{}
	orgService      service.OrganizationServiceInterface = &service.OrganizationService{}
	campaignService service.CampaignServiceInterface     = &service.CampaignService{}
)
