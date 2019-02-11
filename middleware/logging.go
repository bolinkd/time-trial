package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
)

func LoggingMiddleware(c *gin.Context) {
	fields := make(map[string]interface{})
	url := c.Request.URL.Path

	// don't log pings
	if strings.Compare(url, "/") != 0 {
		fields["url"] = url
		if c.Request != nil && c.Request.Body != nil {
			buf, _ := ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))

			if buf != nil {
				body := string(buf)
				if body != "" {
					fields["body"] = body
				}
			}
		}
		//complete handler call so that response code is set
		c.Next()
		fields["response_code"] = c.Writer.Status()
		log.WithFields(fields).Info("client request - traxone-service")
	}
}
