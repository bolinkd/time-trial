package handlers

import (
	"errors"
	"fmt"
	"github.com/bolinkd/time-trial-service/middleware"
	"github.com/gin-gonic/gin"
	"strconv"
)

var ErrForbidden = errors.New("access forbidden")

func GetTimeTrialById(context *gin.Context) {
	database := middleware.GetDatabase(context)
	userIDToken := middleware.GetUserID(context)
	userID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		BadRequest(context, fmt.Sprintf("invalid userID: %d", userID))
		return
	}

	err = CheckAuthorizationWithUserID(database, userID, userIDToken)
	if err != nil {
		if err == ErrForbidden {
			Forbidden(context)
		} else {
			UnexpectedError(context, err)
		}
		return
	}

	user, err := userService.GetUserByID(database, userID)
	if err == service.ErrUserNotFound {
		NotFound(context, err)
	} else if err != nil {
		UnexpectedError(context, err)
	} else {
		Ok(context, user)
	}
}

func GetUser(context *gin.Context) {
	database := middleware.GetDatabase(context)
	userID := middleware.GetUserID(context)

	if userID == middleware.UserNotFound {
		Forbidden(context)
		return
	}
	user, err := userService.GetUserByID(database, userID)
	if err == service.ErrUserNotFound {
		NotFound(context, err)
	} else if err != nil {
		UnexpectedError(context, err)
	} else {
		Ok(context, user)
	}
}

func CreateUser(context *gin.Context) {
	database := middleware.GetDatabase(context)

	var userCreation domain.UserCreation
	err := decodeAndValidate(context, &userCreation)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}
	//create org
	tx, err := database.GetTransactor()
	if err != nil {
		UnexpectedError(context, err)
		return
	}

	org, err := orgService.CreateOrg(database, userCreation.Organization, tx)
	if err != nil {
		UnexpectedError(context, err)
		return
	}

	userCreation.User.OrgID = org.ID
	//create user
	user, err := userService.AddUser(database, userCreation.User, tx)
	if err != nil {
		if err == service.ErrOrgNotFound {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		if tx != nil {
			err = tx.Commit()
			if err != nil {
				UnexpectedError(context, err)
				return
			}
		}
		Created(context, user)

	}
}

func UpdateUser(context *gin.Context) {
	database := middleware.GetDatabase(context)
	userID := middleware.GetUserID(context)

	var userD domain.Traxuser
	err := decodeAndValidate(context, &userD)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	err = CheckAuthorizationWithUserID(database, userD.ID, userID)
	if err != nil {
		if err == ErrForbidden {
			Forbidden(context)
		} else {
			UnexpectedError(context, err)
		}
		return
	}

	user, err := userService.UpdateUser(database, userD)
	if err != nil {
		if err == service.ErrOrgNotFound {
			BadRequest(context, err.Error())
		}
		if err == service.ErrUserNotFound {
			NotFound(context, err)
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Ok(context, user)
	}
}
