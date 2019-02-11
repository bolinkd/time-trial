package middleware

import (
	"github.com/auth0-community/go-auth0"
	"github.com/businessinstincts/traxone/common"
	"github.com/businessinstincts/traxone/domain"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"net/http"
)

var (
	Domain          = common.GetEnv([]string{"AUTH0_DOMAIN"}, "https://level48.auth0.com")
	AudienceMachine = []string{
		common.GetEnv([]string{"AUTH0_AUDIENCE_MACHINE"}, "https://level48.auth0.com/api/v2/"),
	}
)

const (
	UserIDKey    = "user-id"
	userIDClaims = "https://traxone.level48.com/user_id"
)

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

func JWTMachineMiddleware(c *gin.Context) {
	JWKS_URI := Domain + "/.well-known/jwks.json"
	client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: JWKS_URI}, nil)

	var ApiIssuer = Domain + "/"
	configuration := auth0.NewConfiguration(client, AudienceMachine, ApiIssuer, jose.RS256)
	validator := auth0.NewValidator(configuration, nil)

	token, err := validator.ValidateRequest(c.Request)

	if err != nil {
		log.Println("Machine Token is not valid or missing,", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, domain.Message("Missing or invalid token."))
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

func checkScope(r *http.Request, validator *auth0.JWTValidator, token *jwt.JSONWebToken) (map[string]interface{}, error) {
	claims := map[string]interface{}{}
	err := validator.Claims(r, token, &claims)
	return claims, err
}
