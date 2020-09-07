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

// abortStatus aborts the current context and sets response status with and its relative content
func abortStatus(c *gin.Context, httpStatusCode int) {
	c.AbortWithStatusJSON(
		httpStatusCode,
		gin.H{errorResponseKey: http.StatusText(httpStatusCode)},
	)
}

// abortError aborts with generic http status content or the error reason depending on api mode
func abortError(c *gin.Context, httpStatusCode int, err error) {
	apiMode := gin.Mode()
	if apiMode == gin.DebugMode || apiMode == gin.TestMode {
		c.AbortWithStatusJSON(
			httpStatusCode,
			gin.H{errorResponseKey: err.Error()},
		)
	} else {
		abortStatus(c, httpStatusCode)
	}
}
