package middleware

import (
	"github.com/auth0-community/go-auth0"
	"github.com/businessinstincts/traxone/common"
	"github.com/businessinstincts/traxone/domain"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/square/go-jose.v2"
	"net/http"
)

var (
	AudienceSPA = []string{
		common.GetEnv([]string{"AUTH0_AUDIENCE_SPA"}, "lAe4fAYpsX3J9FIv3i1w0aLUbyxgLKoH"),
	}
)

const (
	UserNotFound = 0
	UserLocal    = 1
)

func JWTSPAMiddleware(c *gin.Context) {
	JWKS_URI := Domain + "/.well-known/jwks.json"
	client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: JWKS_URI}, nil)

	var ApiIssuer = Domain + "/"
	configuration := auth0.NewConfiguration(client, AudienceSPA, ApiIssuer, jose.RS256)
	validator := auth0.NewValidator(configuration, nil)

	token, err := validator.ValidateRequest(c.Request)
	if err != nil {
		if common.IsPhaseLocal() {
			c.Set(UserIDKey, UserLocal)
			c.Next()
		} else {
			log.Println("SPA Token is not valid or missing,", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.Message("Missing or invalid token."))
		}
		return
	}
	// Ensure the token has the correct scope
	claims, err := checkScope(c.Request, validator, token)
	if err != nil {
		// If the token is valid and we have the right scope, we'll pass through the middleware
		c.AbortWithStatusJSON(http.StatusForbidden, domain.Message("You do not have the read:messages scope."))
		return
	}
	if claims[userIDClaims] == nil {
		if err != nil {
			log.Error(err)
		}
	} else {
		c.Set(UserIDKey, int(claims[userIDClaims].(float64)))
	}
	c.Next()
}

func GetUserID(context *gin.Context) int {
	userID, exists := context.Keys[UserIDKey]
	if exists {
		return userID.(int)
	} else {
		return UserNotFound
	}
}
