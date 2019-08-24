package handlers

import (
	"fmt"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/middleware"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetRentalRowersByRental(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	rentalID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		BadRequest(context, fmt.Sprintf("invalid id: %d", rentalID))
		return
	}

	rentalRowers, err := services.GetRentalRowersByRental(database, rentalID)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Ok(context, domain.RentalRowerSlice{rentalRowers})
	}
}

func GetRentalRowerByID(context *gin.Context) {
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

func CreateRentalRower(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	var rentalRower domain.RentalRower
	err := decodeAndValidate(context, &rentalRower)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	err = services.CreateRentalRower(database, rentalRower.RentalRower)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Created(context, rentalRower)
	}
}

func UpdateRentalRower(context *gin.Context) {
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

func DeleteRentalRower(context *gin.Context) {
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
