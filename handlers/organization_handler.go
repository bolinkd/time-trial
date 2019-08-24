package handlers

import (
	"fmt"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/middleware"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetOrganizations(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	orgs, err := services.GetOrganizations(database)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	}
	Ok(context, domain.OrganizationSlice{orgs})
}

func GetOrganizationByID(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	orgID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		BadRequest(context, fmt.Sprintf("invalid id: %d", orgID))
		return
	}

	org, err := services.GetOrganizationByID(database, orgID)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Ok(context, &domain.Organization{org})
	}
}

func CreateOrganization(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	var org domain.Organization
	err := decodeAndValidate(context, &org)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	err = services.CreateOrganization(database, org.Organization)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Created(context, org)
	}
}

func UpdateOrganization(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	var org domain.Organization
	err := decodeAndValidate(context, &org)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	err = services.UpdateOrganization(database, org.Organization)
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
