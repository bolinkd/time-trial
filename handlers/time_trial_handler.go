package handlers

import (
	"errors"
	"fmt"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/middleware"
	"github.com/bolinkd/time-trial/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

var ErrForbidden = errors.New("access forbidden")

func GetTimeTrials(context *gin.Context) {
	database := middleware.GetDatabase(context)

	user, err := timeTrialService.GetTimeTrials(database)
	if err == service.ErrTimeTrialNotFound {
		NotFound(context, err)
	} else if err != nil {
		UnexpectedError(context, err)
	} else {
		Ok(context, user)
	}
}

func GetTimeTrialById(context *gin.Context) {
	database := middleware.GetDatabase(context)

	timeTrialID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		BadRequest(context, fmt.Sprintf("invalid userID: %d", timeTrialID))
		return
	}

	timeTrial, err := timeTrialService.GetTimeTrialById(database, timeTrialID)
	if err == service.ErrTimeTrialNotFound {
		NotFound(context, err)
	} else if err != nil {
		UnexpectedError(context, err)
	} else {
		Ok(context, timeTrial)
	}
}

func CreateTimeTrial(context *gin.Context) {
	database := middleware.GetDatabase(context)

	var timeTrialD domain.TimeTrial
	err := decodeAndValidate(context, &timeTrialD)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	timeTrial, err := timeTrialService.CreateTimeTrial(database, timeTrialD, nil)
	if err != nil {
		if _, ok := err.(domain.TraxError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Created(context, timeTrial)
	}
}

func UpdateTimeTrial(context *gin.Context) {
	database := middleware.GetDatabase(context)

	var timeTrialD domain.TimeTrial
	err := decodeAndValidate(context, &timeTrialD)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	timeTrial, err := timeTrialService.UpdateTimeTrial(database, timeTrialD)
	if err != nil {
		if err == service.ErrTimeTrialNotFound {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Ok(context, timeTrial)
	}
}
