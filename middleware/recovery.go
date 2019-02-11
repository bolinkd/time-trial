package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"runtime/debug"
)

func RecoveryHandler(c *gin.Context) {
	var err error
	defer func() {
		r := recover()
		if r != nil {
			switch t := r.(type) {
			case string:
				err = errors.New(t)
				break
			case error:
				err = t
				break
			default:
				err = errors.New(fmt.Sprint("unknown error: ", r))
				break
			}

			logrus.WithError(err).WithField("stacktrace", string(debug.Stack())).Error()
			err := c.AbortWithError(http.StatusInternalServerError, err)
			if err != nil {
				log.Println(err)
			}
		}
	}()
	c.Next()
}
