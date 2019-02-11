package middleware

import (
	"fmt"
	"github.com/businessinstincts/traxone/db"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	DbServerKey = "db"
)

func (s *Server) DbHandler(context *gin.Context) {
	if _, ok := context.Keys[DbServerKey]; !ok {
		context.Set(DbServerKey, s.Database)
	}
}

func retry(attempts int, sleep time.Duration, callback func() error) (err error) {
	for i := 0; ; i++ {
		err = callback()
		if err == nil {
			return
		}

		if i >= (attempts - 1) {
			break
		}

		time.Sleep(sleep)
		fmt.Println("retrying database connection after error:", err)
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

func GetDatabase(context *gin.Context) db.DatabaseInterface {
	return context.Keys[DbServerKey].(db.DatabaseInterface)
}
