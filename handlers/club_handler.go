package handlers

import (
	"fmt"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/middleware"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetClubsByOrganization(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	clubID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		BadRequest(context, fmt.Sprintf("invalid id: %d", clubID))
		return
	}

	clubs, err := services.GetClubsByOrganization(database, clubID)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	}
	Ok(context, clubs)
}

func GetClubByID(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	clubID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		BadRequest(context, fmt.Sprintf("invalid id: %d", clubID))
		return
	}

	org, err := services.GetClubByID(database, clubID)
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

func CreateClub(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	var club domain.Club
	err := decodeAndValidate(context, &club)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	err = services.CreateClub(database, club.Club)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Created(context, club)
	}
}

func UpdateClub(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	var club domain.Club
	err := decodeAndValidate(context, &club)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	err = services.UpdateClub(database, club.Club)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Ok(context, club)
	}
}
