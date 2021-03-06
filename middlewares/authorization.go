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

// JWT passphrase
var jwtPassphrase []byte

// SetAuthorizationPassphrase set the passphrase to encrypt/decrypt tokens for JWT
func SetAuthorizationPassphrase(passphrase string) {
	jwtPassphrase = []byte(passphrase)
}

// CustomClaims reprensents our custom jwt claims
type CustomClaims struct {
	jwt.StandardClaims
	UserID uint
}

// CreateAuthToken generates an authorization token
func CreateAuthToken(userID uint) (authToken string, err error) {
	expirationTime := time.Now().Add(authenticationDuration)

	claims := &CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	authToken, err = token.SignedString(jwtPassphrase)
	return
}

// Authorization wraps over endpoint requiring authorization
func Authorization(c *gin.Context) {
	authorizationHeader := c.GetHeader(AuthorizationHeader)
	authToken := strings.Replace(authorizationHeader, AuthorizationHeaderType+" ", "", 1)

	claims := &CustomClaims{}
	token, err := jwt.ParseWithClaims(authToken, claims, func(tkn *jwt.Token) (interface{}, error) {
		return jwtPassphrase, nil
	})

	// Authorization granted
	if err == nil && token.Valid {
		c.Next()
		return
	}

	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": http.StatusText(http.StatusUnauthorized),
	})
}
