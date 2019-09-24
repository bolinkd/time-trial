package handlers

import (
	"fmt"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/middleware"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetRowersByGroup(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	groupID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		BadRequest(context, fmt.Sprintf("invalid id: %d", groupID))
		return
	}

	rowers, err := services.GetRowersByGroup(database, groupID)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	}
	Ok(context, rowers)
}

func GetRowersByCurrentOrganization(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	orgID, err := getCurrentOrganizationID(context)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	rowers, err := services.GetRowersByOrganization(database, orgID)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
		return
	}
	Ok(context, domain.RowerSlice{
		RowerSlice: rowers,
	})
}

func GetRowerByID(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	rowerID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		BadRequest(context, fmt.Sprintf("invalid id: %d", rowerID))
		return
	}

	org, err := services.GetRowerByID(database, rowerID)
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

func CreateRower(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	var rower domain.Rower
	err := decodeAndValidate(context, &rower)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	err = services.CreateRower(database, rower.Rower)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Created(context, rower)
	}
}

func UpdateRower(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	var rower domain.Rower
	err := decodeAndValidate(context, &rower)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	err = services.UpdateRower(database, rower.Rower)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Ok(context, rower)
	}
}
