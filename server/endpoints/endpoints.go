package endpoints

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github/vriaan/footballmanagerapi/server/endpoints/handlers"
)

// List of all endpoints managed by the API, new endpoints declaration must be added here
var endpoints = []Endpoint{
	Endpoint{http.MethodGet, "/ping", handlers.Ping},
	Endpoint{http.MethodPost, "/footballers", handlers.RegisterNewFootballer},
	Endpoint{http.MethodGet, "/footballers/:id", handlers.GetFootballer},
	Endpoint{http.MethodGet, "/footballers", handlers.GetFootballers},
	Endpoint{http.MethodPost, "/login", handlers.Login},
}

// Endpoint represents an API Endpoint
type Endpoint struct {
	Method   string
	Path     string
	Function func(c *gin.Context)
}

// Get returns a copy of API declared endpoint
func Get() (endpointsCopy []Endpoint) {
	copy(endpointsCopy, endpoints)
	return
}

// Register declares all routes managed by the API
func Register(apiEngine *gin.Engine) (err error) {
	for _, endpoint := range endpoints {
		switch endpoint.Method {
		case http.MethodGet:
			apiEngine.GET(endpoint.Path, endpoint.Function)
		case http.MethodPut:
			apiEngine.PUT(endpoint.Path, endpoint.Function)
		case http.MethodPost:
			apiEngine.POST(endpoint.Path, endpoint.Function)
		case http.MethodDelete:
			apiEngine.DELETE(endpoint.Path, endpoint.Function)
		default:
			err = fmt.Errorf("Endpoint %s uses an unknown method %s", endpoint.Path, endpoint.Method)
			return
		}
	}

	return
}
