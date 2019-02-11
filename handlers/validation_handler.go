package handlers

import (
	"encoding/json"
	"github.com/businessinstincts/traxone/domain"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"io/ioutil"
)

var (
	ErrInvalidBody = errors.New("failed to parse body")
)

func decodeAndValidate(c *gin.Context, v domain.InputValidation) error {
	if c.Request.Body == nil {
		return ErrInvalidBody
	}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	if len(body) == 0 {
		return ErrInvalidBody
	}

	err = json.Unmarshal(body, &v)
	if err != nil {
		return err
	}
	return v.Validate()
}
