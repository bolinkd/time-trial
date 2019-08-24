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

func GetBoatsForTimeTrial(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	timeTrialID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		BadRequest(context, fmt.Sprintf("invalid id: %d", timeTrialID))
		return
	}

	user, err := services.GetBoatsForTimeTrial(database, timeTrialID)
	if err != nil {
		UnexpectedError(context, err)
	} else {
		Ok(context, user)
	}
}

func GetBoatByID(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	boatID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		BadRequest(context, fmt.Sprintf("invalid id: %d", boatID))
		return
	}

	timeTrial, err := services.GetBoatByID(database, boatID)
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

func CreateBoat(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)
	socketClient := middleware.GetSocket(context)

	var boat domain.Boat
	err := decodeAndValidate(context, &boat)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	err = services.CreateBoat(database, boat.Boat)
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
				DataType:     "boat",
				Payload:      boat,
			},
		})
		Created(context, boat)
	}
}

func UpdateBoat(context *gin.Context) {
	database := middleware.GetDatabase(context)
	socketClient := middleware.GetSocket(context)
	services := middleware.GetServices(context)

	var boat domain.Boat
	err := decodeAndValidate(context, &boat)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	err = services.UpdateBoat(database, boat.Boat)
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
				DataType:     "boat",
				Payload:      boat,
			},
		})
		Ok(context, boat)
	}
}
