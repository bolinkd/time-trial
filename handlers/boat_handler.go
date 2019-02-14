package handlers

import (
	"fmt"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/middleware"
	"github.com/bolinkd/time-trial/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetBoatsForTimeTrial(context *gin.Context) {
	database := middleware.GetDatabase(context)

	timeTrialID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		BadRequest(context, fmt.Sprintf("invalid id: %d", timeTrialID))
		return
	}

	user, err := boatService.GetBoatsForTimeTrial(database, timeTrialID)
	if err != nil {
		UnexpectedError(context, err)
	} else {
		Ok(context, user)
	}
}

func GetBoatByID(context *gin.Context) {
	database := middleware.GetDatabase(context)

	boatID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		BadRequest(context, fmt.Sprintf("invalid id: %d", boatID))
		return
	}

	timeTrial, err := boatService.GetBoatByID(database, boatID)
	if err == service.ErrBoatNotFound {
		NotFound(context, err)
	} else if err != nil {
		UnexpectedError(context, err)
	} else {
		Ok(context, timeTrial)
	}
}

func CreateBoat(context *gin.Context) {
	database := middleware.GetDatabase(context)

	var boatD domain.Boat
	err := decodeAndValidate(context, &boatD)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	boat, err := boatService.CreateBoat(database, boatD)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Created(context, boat)
	}
}

func UpdateBoat(context *gin.Context) {
	database := middleware.GetDatabase(context)

	var boatD domain.Boat
	err := decodeAndValidate(context, &boatD)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	boat, err := boatService.UpdateBoat(database, boatD)
	if err != nil {
		if err == service.ErrBoatNotFound {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Ok(context, boat)
	}
}
