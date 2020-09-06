package middlewares

import (
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	// Token validity duration
	authenticationDuration = 1 * time.Hour
	// AuthorizationHeaderType is the authorization Header type
	AuthorizationHeaderType = "Bearer"
	// AuthorizationHeader is the header name for authorization
	AuthorizationHeader = "Authorization"
)

// TODO: move this within a file, env setting or something else
var jwtSecretKey = []byte("thisissecret")

// CustomClaims reprensents our custom jwt claims
type CustomClaims struct {
	jwt.StandardClaims
	UserID uint
}

// AttachAuthenticationToken generates and set header for authorization
func AttachAuthenticationToken(c *gin.Context, userID uint) error {
	authToken, err := createAuthToken(userID)
	if err == nil {
		c.Header(AuthorizationHeader, AuthorizationHeaderType+" "+authToken)
	}
	return err
}

// createAuthToken generates a JWT token based on secret key
func createAuthToken(userID uint) (jwtAuthToken string, err error) {
	expirationTime := time.Now().Add(authenticationDuration)

	claims := &CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtAuthToken, err = token.SignedString(jwtSecretKey)
	return
}

// Authorization wraps over endpoint requiring authorization
func Authorization(c *gin.Context) {
	authorizationHeader := c.GetHeader(AuthorizationHeader)
	authToken := strings.Replace(authorizationHeader, AuthorizationHeaderType+" ", "", 1)

	claims := CustomClaims{}
	token, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": http.StatusText(http.StatusUnauthorized),
		})
		// return
	}
}
