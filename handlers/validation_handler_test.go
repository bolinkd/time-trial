package handlers

import (
	"bytes"
	"github.com/businessinstincts/traxone/domain"
	"github.com/businessinstincts/traxone/utils/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_decodeAndValidate(t *testing.T) {
	r.POST("/", func(context *gin.Context) {
		var v domain.FeatureCollection
		err := decodeAndValidate(context, &v)
		if err != nil {
			context.JSON(http.StatusBadRequest, domain.Message(err.Error()))
		} else {
			context.JSON(http.StatusOK, nil)
		}
	})

	t.Run("Empty Body", func(t *testing.T) {
		w, err := makeRequest(http.MethodPost, "/", mocks.UserStandard1, nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)

		w, err = makeRequest(http.MethodPost, "/", mocks.UserStandard1, bytes.NewReader([]byte("")))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Valid", func(t *testing.T) {
		body := []byte(`{"type":"FeatureCollection","features":[{"type":"Feature","geometry":{"type":"PointAndRadius","point":{"lat":30.30676,"lng":-97.71065899999996},"radius":200},"properties":{"start":"2018-08-30T23:59","end":"2018-11-29T23:59"}},{"type":"Feature","geometry":{"type":"PointAndRadius","point":{"lat":30.15923799999999,"lng":-97.79185100000001},"radius":200},"properties":{"start":"2018-08-30T23:59","end":"2018-11-29T23:59"}},{"type":"Feature","geometry":{"type":"PointAndRadius","point":{"lat":30.468736,"lng":-97.80060200000003},"radius":200},"properties":{"start":"2018-08-30T23:59","end":"2018-11-29T23:59"}},{"type":"Feature","geometry":{"type":"PointAndRadius","point":{"lat":30.473232,"lng":-97.59680600000002},"radius":200},"properties":{"start":"2018-08-30T23:59","end":"2018-11-29T23:59"}},{"type":"Feature","geometry":{"type":"PointAndRadius","point":{"lat":30.63005,"lng":-97.69346389999998},"radius":200},"properties":{"start":"2018-08-30T23:59","end":"2018-11-29T23:59"}}]}`)

		w, err := makeRequest(http.MethodPost, "/", mocks.UserStandard1, bytes.NewReader(body))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)

	})
}
