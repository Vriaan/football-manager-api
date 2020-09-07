package actions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	// PaginationSizeKey is page result size query parameter name
	PaginationSizeKey = "limit"
	// DefaultPaginationSize is page result size per default
	DefaultPaginationSize = "100"
	// PaginationOffsetKey is page result offset query parameter name
	PaginationOffsetKey = "offset"
	// DefaultOffset is page result offset per default
	DefaultOffset = "0"

	errorResponseKey = "error"
)

// Action represents the endpoint Action to be called (same as the controllerAction)
type Action gin.HandlerFunc

// Paginate gets a database handler and adds the limit & offset from asked from request or set default
func paginate(c *gin.Context) (limit, offset string) {
	limit = c.DefaultQuery(PaginationSizeKey, DefaultPaginationSize)
	offset = c.DefaultQuery(PaginationOffsetKey, DefaultOffset)
	return
}

func writeStatusError(c *gin.Context, httpStatusCode int) {
	c.AbortWithStatusJSON(
		httpStatusCode,
		gin.H{"error": http.StatusText(httpStatusCode)},
	)
}

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
