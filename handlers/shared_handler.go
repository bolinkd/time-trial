package handlers

import (
	"github.com/bolinkd/time-trial/service"
)

var (
	timeTrialService service.TimeTrialServiceInterface = &service.TimeTrialService{}
	boatService      service.BoatServiceInterface      = &service.BoatService{}
)
