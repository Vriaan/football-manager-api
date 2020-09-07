package actions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is the route to call to test API is alive
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
