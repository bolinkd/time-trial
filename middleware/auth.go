package middleware

import (
	"fmt"
	"github.com/bolinkd/time-trial/db"
	"github.com/bolinkd/time-trial/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AuthMiddleware(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	orgID, err := strconv.Atoi(c.Request.Header.Get("Organization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, domain.Message("Unauthorized - unable to parse orgID"))
		return
	}

	fmt.Println(token)
	database := GetDatabase(c).(db.AuthDBInterface)

	valid, err := database.ValidateAuthToken(nil, orgID, token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, domain.Message("Unauthorized - error retrieving auth token"))
		return
	}

	if !valid {
		c.AbortWithStatusJSON(http.StatusForbidden, domain.Message("Unauthorized - invalid auth token"))
		return
	}

	c.Next()
}
