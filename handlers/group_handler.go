package handlers

import (
	"fmt"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/middleware"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetGroupsByClub(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	clubID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		BadRequest(context, fmt.Sprintf("invalid id: %d", clubID))
		return
	}

	groups, err := services.GetGroupsByClub(database, clubID)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	}
	Ok(context, groups)
}

func GetGroupByID(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	groupID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		BadRequest(context, fmt.Sprintf("invalid id: %d", groupID))
		return
	}

	org, err := services.GetGroupByID(database, groupID)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Ok(context, org)
	}
}

func CreateGroup(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	var group domain.Group
	err := decodeAndValidate(context, &group)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	err = services.CreateGroup(database, group.Group)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Created(context, group)
	}
}

func UpdateGroup(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	var group domain.Group
	err := decodeAndValidate(context, &group)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	err = services.UpdateGroup(database, group.Group)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Ok(context, group)
	}
}
