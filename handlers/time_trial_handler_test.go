package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/businessinstincts/traxone/domain"
	"github.com/businessinstincts/traxone/utils/mocks"
	"github.com/stretchr/testify/assert"
	"gopkg.in/volatiletech/null.v6"
	"net/http"
	"testing"
)

var (
	r = mocks.NewGinMock()
)

func TestGetUserByID(t *testing.T) {
	r.GET("/v1/users/:id", GetUserByID)
	// Create the mock request you'd like to test. Make sure the second argument
	// here is the same as one of the routes you defined in the router setup
	// block!

	t.Run("User Not Found", func(t *testing.T) {
		w, err := makeRequest(http.MethodGet, fmt.Sprintf("/v1/users/%d", mocks.UserNotFound), mocks.UserAdmin, nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Invalid User ID", func(t *testing.T) {
		w, err := makeRequest(http.MethodGet, fmt.Sprintf("/v1/users/%s", "test"), mocks.UserStandard1, nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Valid User ID - Standard", func(t *testing.T) {
		w, err := makeRequest(http.MethodGet, fmt.Sprintf("/v1/users/%d", mocks.UserStandard1), mocks.UserStandard1, nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Forbidden", func(t *testing.T) {
		w, err := makeRequest(http.MethodGet, fmt.Sprintf("/v1/users/%d", mocks.UserStandard1), mocks.UserStandard2, nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("Valid User ID - Admin", func(t *testing.T) {
		w, err := makeRequest(http.MethodGet, fmt.Sprintf("/v1/users/%d", mocks.UserStandard1), mocks.UserAdmin, nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Unexpected Error", func(t *testing.T) {
		w, err := makeRequest(http.MethodGet, fmt.Sprintf("/v1/users/%d", mocks.UserUnknownError), mocks.UserAdmin, nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

}

func TestGetUser(t *testing.T) {
	r.GET("/v1/users", GetUser)

	t.Run("Valid Token", func(t *testing.T) {
		w, err := makeRequest(http.MethodGet, "/v1/users", mocks.UserStandard1, nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestCreateUser(t *testing.T) {
	r.POST("/v1/users", CreateUser)

	t.Run("Bad Body", func(t *testing.T) {
		bodyBytes, err := json.Marshal("")
		assert.NoError(t, err)

		body := bytes.NewReader(bodyBytes)

		w, err := makeRequest(http.MethodPost, "/v1/users", mocks.UserStandard1, body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Fail to create Organization", func(t *testing.T) {
		userCreate := domain.UserCreation{
			User: domain.Traxuser{
				FirstName: null.StringFrom("Leeroy"),
				LastName:  null.StringFrom("Jenkins"),
				Email:     "leeroy@jenkins.com",
				Role:      0,
			},
			Organization: domain.Organization{
				ID:          mocks.OrgUnknownError,
				Description: "test org",
			},
		}

		bodyBytes, err := json.Marshal(userCreate)
		assert.NoError(t, err)

		body := bytes.NewReader(bodyBytes)

		w, err := makeRequest(http.MethodPost, "/v1/users", mocks.UserStandard1, body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("Fail to create User", func(t *testing.T) {
		userCreate := domain.UserCreation{
			User: domain.Traxuser{
				ID:        mocks.UserUnknownError,
				FirstName: null.StringFrom("Leeroy"),
				LastName:  null.StringFrom("Jenkins"),
				Email:     "leeroy@jenkins.com",
				Role:      0,
			},
			Organization: domain.Organization{
				Description: "test org",
			},
		}

		bodyBytes, err := json.Marshal(userCreate)
		assert.NoError(t, err)

		body := bytes.NewReader(bodyBytes)

		w, err := makeRequest(http.MethodPost, "/v1/users", mocks.UserStandard1, body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		userCreate := domain.UserCreation{
			User: domain.Traxuser{
				FirstName: null.StringFrom("Leeroy"),
				LastName:  null.StringFrom("Jenkins"),
				Email:     "leeroy@jenkins.com",
				Role:      0,
			},
			Organization: domain.Organization{
				Description: "test org",
			},
		}

		bodyBytes, err := json.Marshal(userCreate)
		assert.NoError(t, err)

		body := bytes.NewReader(bodyBytes)

		w, err := makeRequest(http.MethodPost, "/v1/users", mocks.UserStandard1, body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, w.Code)
	})
}

func TestUpdateUser(t *testing.T) {
	r.PUT("/v1/users", UpdateUser)

	t.Run("Bad Body", func(t *testing.T) {

		bodyBytes, err := json.Marshal("")
		assert.NoError(t, err)

		body := bytes.NewReader(bodyBytes)

		w, err := makeRequest(http.MethodPut, "/v1/users", mocks.UserStandard1, body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		user := domain.Traxuser{
			ID:        mocks.UserStandard1,
			OrgID:     mocks.OrgStandard1,
			FirstName: null.StringFrom("Leeroy"),
			LastName:  null.StringFrom("Jenkins"),
			Email:     "leeroy@jenkins.com",
			Role:      0,
		}

		bodyBytes, err := json.Marshal(user)
		assert.NoError(t, err)

		body := bytes.NewReader(bodyBytes)

		w, err := makeRequest(http.MethodPut, "/v1/users", mocks.UserStandard2, body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("User Not Found", func(t *testing.T) {
		user := domain.Traxuser{
			ID:        mocks.UserNotFound,
			OrgID:     mocks.OrgNotFound,
			FirstName: null.StringFrom("Leeroy"),
			LastName:  null.StringFrom("Jenkins"),
			Email:     "leeroy@jenkins.com",
			Role:      0,
		}

		bodyBytes, err := json.Marshal(user)
		assert.NoError(t, err)

		body := bytes.NewReader(bodyBytes)

		w, err := makeRequest(http.MethodPut, "/v1/users", mocks.UserAdmin, body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Org Not Found", func(t *testing.T) {
		user := domain.Traxuser{
			ID:        mocks.UserStandard1,
			OrgID:     mocks.OrgNotFound,
			FirstName: null.StringFrom("Leeroy"),
			LastName:  null.StringFrom("Jenkins"),
			Email:     "leeroy@jenkins.com",
			Role:      0,
		}

		bodyBytes, err := json.Marshal(user)
		assert.NoError(t, err)

		body := bytes.NewReader(bodyBytes)

		w, err := makeRequest(http.MethodPut, "/v1/users", mocks.UserStandard1, body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		user := domain.Traxuser{
			ID:        mocks.UserStandard1,
			OrgID:     mocks.OrgStandard1,
			FirstName: null.StringFrom("Leeroy"),
			LastName:  null.StringFrom("Jenkins"),
			Email:     "leeroy@jenkins.com",
			Role:      0,
		}

		bodyBytes, err := json.Marshal(user)
		assert.NoError(t, err)

		body := bytes.NewReader(bodyBytes)

		w, err := makeRequest(http.MethodPut, "/v1/users", mocks.UserStandard1, body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
