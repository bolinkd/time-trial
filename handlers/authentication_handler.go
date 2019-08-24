package handlers

import (
	"github.com/bolinkd/time-trial/domain"
	"github.com/bolinkd/time-trial/middleware"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	database := middleware.GetDatabase(context)

	var orgAuth domain.OrganizationAuth
	err := decodeAndValidate(context, &orgAuth)
	if err != nil {
		BadRequest(context, err.Error())
		return
	}

	resp, err := database.GetToken(nil, orgAuth.OrganizationAuth)
	if err != nil {
		if _, ok := err.(domain.AppError); ok {
			BadRequest(context, err.Error())
		} else {
			UnexpectedError(context, err)
		}
	} else {
		Ok(context, resp)
	}
}
