package handlers

import (
	"fmt"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/middleware"
	"github.com/bolinkd/time-trial/socket"
	"github.com/bolinkd/time-trial/socket/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetTimeTrials(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	user, err := services.GetTimeTrials(database)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Ok(context, user)
	}
}

func GetTimeTrialById(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	timeTrialID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		BadRequest(context, fmt.Sprintf("invalid userID: %d", timeTrialID))
		return
	}

	timeTrial, err := services.GetTimeTrialById(database, timeTrialID)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Ok(context, timeTrial)
	}
}

func CreateTimeTrial(context *gin.Context) {
	database := middleware.GetDatabase(context)
	socketClient := middleware.GetSocket(context)
	services := middleware.GetServices(context)

	var timeTrial domain.TimeTrial
	err := decodeAndValidate(context, &timeTrial)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	err = services.CreateTimeTrial(database, timeTrial.TimeTrial)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		go socketClient.SendUpdateEventByRoom(socket.RoomRaceData, model.UpdateEvent{
			Type: "race-data-update",
			Payload: model.Event{
				UpdateMethod: "create",
				DataType:     "time-trial",
				Payload:      timeTrial,
			},
		})
		Created(context, timeTrial)
	}
}

func UpdateTimeTrial(context *gin.Context) {
	database := middleware.GetDatabase(context)
	socketClient := middleware.GetSocket(context)
	services := middleware.GetServices(context)

	var timeTrial domain.TimeTrial
	err := decodeAndValidate(context, &timeTrial)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	err = services.UpdateTimeTrial(database, timeTrial.TimeTrial)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		go socketClient.SendUpdateEventByRoom(socket.RoomRaceData, model.UpdateEvent{
			Type: "race-data-update",
			Payload: model.Event{
				UpdateMethod: "update",
				DataType:     "time-trial",
				Payload:      timeTrial,
			},
		})
		Ok(context, timeTrial)
	}
}
