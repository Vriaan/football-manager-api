package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//TODO could helper be moved to middleware ?

// SetErrorFromStatus returns a pre formatted error message based on current API mode
func SetErrorFromStatus(httpStatusCode int, err error) gin.H {
	var errorMessage string
	switch gin.Mode() {
	case gin.DebugMode:
		fallthrough
	case gin.TestMode:
		errorMessage = err.Error()
	case gin.ReleaseMode:
		fallthrough
	default:
		errorMessage = http.StatusText(httpStatusCode)
	}
	return gin.H{"error": errorMessage}
}
