package handlers

import (
	"github.com/bolinkd/time-trial/domain"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func UnexpectedError(context *gin.Context, err error) {
	log.WithField("error", err.Error()).Info("Unexpected Error")
	context.JSON(http.StatusInternalServerError, domain.Message(err.Error()))
}

func NotFound(context *gin.Context, err error) {
	context.JSON(http.StatusNotFound, domain.Message(err.Error()))
}

func Created(context *gin.Context, item interface{}) {
	context.JSON(http.StatusCreated, item)
}

func Ok(context *gin.Context, item interface{}) {
	context.JSON(http.StatusOK, item)
}

func BadRequest(context *gin.Context, str string) {
	context.JSON(http.StatusBadRequest, domain.Message(str))
}

func Forbidden(context *gin.Context) {
	context.JSON(http.StatusForbidden, domain.Message("Unauthorized"))
}
