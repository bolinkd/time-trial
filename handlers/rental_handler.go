package handlers

import (
	"fmt"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/middleware"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func GetRentals(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	orgID, err := strconv.Atoi(context.Query("org_id"))
	if err != nil {
		BadRequest(context, "invalid organization id")
		return
	}

	active, err := strconv.ParseBool(context.DefaultQuery("active", "false"))
	if err != nil {
		active = false
	}

	startTime, err := time.Parse(time.RFC3339, context.Query("start_time"))
	if err != nil {
		startTime = time.Unix(0, 0).UTC()
	}

	endTime, err := time.Parse(time.RFC3339, context.Query("end_time"))
	if err != nil {
		endTime = time.Now().UTC()
	}

	if startTime.After(endTime) {
		BadRequest(context, "start_time must be before end_time")
		return
	}

	rentals, err := services.GetRentals(database, orgID, active, startTime, endTime)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Ok(context, rentals)
	}
}

func GetRentalsByShell(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	shellID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		BadRequest(context, fmt.Sprintf("invalid id: %d", shellID))
		return
	}

	rentals, err := services.GetRentalsByShell(database, shellID)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	}
	Ok(context, rentals)
}

func GetRentalByID(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	rentalID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		BadRequest(context, fmt.Sprintf("invalid id: %d", rentalID))
		return
	}

	rental, err := services.GetRentalByID(database, rentalID)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Ok(context, rental)
	}
}

func CreateRental(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	var rental domain.Rental
	err := decodeAndValidate(context, &rental)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	err = services.CreateRental(database, rental.Rental)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Created(context, rental)
	}
}

func UpdateRental(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	var rental domain.Rental
	err := decodeAndValidate(context, &rental)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	err = services.UpdateRental(database, rental.Rental)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Ok(context, rental)
	}
}

func DeleteRental(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	rentalID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		BadRequest(context, fmt.Sprintf("invalid id: %d", rentalID))
		return
	}

	err = services.DeleteRental(database, rentalID)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Ok(context, "deleted")
	}
}
