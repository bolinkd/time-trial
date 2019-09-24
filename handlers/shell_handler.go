package handlers

import (
	"fmt"
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/middleware"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetShellsByCurrentOrganization(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	orgID, err := getCurrentOrganizationID(context)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	shells, err := services.GetShellsByOrganization(database, orgID)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	}
	Ok(context, domain.ShellSlice{
		ShellSlice: shells,
	})
}

func GetShellsByGroup(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	groupID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		BadRequest(context, fmt.Sprintf("invalid id: %d", groupID))
		return
	}

	shells, err := services.GetShellsByGroup(database, groupID)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	}
	Ok(context, domain.ShellSlice{
		ShellSlice: shells,
	})
}

func GetShellByID(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	shellID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		BadRequest(context, fmt.Sprintf("invalid id: %d", shellID))
		return
	}

	shell, err := services.GetShellByID(database, shellID)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Ok(context, shell)
	}
}

func CreateShell(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	var shell domain.Shell
	err := decodeAndValidate(context, &shell)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	err = services.CreateShell(database, shell.Shell)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Created(context, shell)
	}
}

func UpdateShell(context *gin.Context) {
	database := middleware.GetDatabase(context)
	services := middleware.GetServices(context)

	var shell domain.Shell
	err := decodeAndValidate(context, &shell)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	err = services.UpdateShell(database, shell.Shell)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Ok(context, shell)
	}
}
