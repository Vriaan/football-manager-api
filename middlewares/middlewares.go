package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Middleware represents a middleware to be called before the actions
type Middleware gin.HandlerFunc

// Middlewares represents a list of middleware
type Middlewares []gin.HandlerFunc
