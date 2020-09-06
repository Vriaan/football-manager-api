package endpoints

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github/vriaan/footballmanagerapi/server/endpoints/handlers"
	"github/vriaan/footballmanagerapi/server/middlewares"
)

// List of all endpoints managed by the API, new endpoints declaration must be added here
var (
	// Those listed endpoint check for authorization before being accessed
	needAuthorizationEndpoints = gin.RoutesInfo{
		gin.RouteInfo{Method: http.MethodPost, Path: "/footballers", HandlerFunc: handlers.RegisterNewFootballer},
		gin.RouteInfo{Method: http.MethodGet, Path: "/footballers", HandlerFunc: handlers.ListFootballers},
		gin.RouteInfo{Method: http.MethodGet, Path: "/footballers/:id", HandlerFunc: handlers.GetFootballer},
		gin.RouteInfo{Method: http.MethodDelete, Path: "/footballers/:id", HandlerFunc: handlers.DeleteFootballer},
		gin.RouteInfo{Method: http.MethodPut, Path: "/footballers/:id", HandlerFunc: handlers.UpdateFootballer},
	}

	// Those listed endpoints access is not protected by authorization
	freeAccessEndpoints = gin.RoutesInfo{
		gin.RouteInfo{Method: http.MethodGet, Path: "/ping", HandlerFunc: handlers.Ping},
		gin.RouteInfo{Method: http.MethodPost, Path: "/login", HandlerFunc: handlers.Login},
	}
)

// Get returns a copy of API declared endpoint
func Get() (endpointsCopy gin.RoutesInfo) {
	endpointsCopy = append(freeAccessEndpoints, needAuthorizationEndpoints...)
	// endpointsCopy = make([]gin.RouteInfo, len(endpoints))
	// copy(endpointsCopy, endpoints)
	return
}

// Register declares all routes managed by the API
func Register(apiEngine *gin.Engine) (err error) {
	for _, endpoint := range freeAccessEndpoints {
		switch endpoint.Method {
		case http.MethodGet:
			apiEngine.GET(endpoint.Path, endpoint.HandlerFunc)
		case http.MethodPut:
			apiEngine.PUT(endpoint.Path, endpoint.HandlerFunc)
		case http.MethodPost:
			apiEngine.POST(endpoint.Path, endpoint.HandlerFunc)
		case http.MethodDelete:
			apiEngine.DELETE(endpoint.Path, endpoint.HandlerFunc)
		default:
			err = fmt.Errorf("Endpoint %s uses an unknown method %s", endpoint.Path, endpoint.Method)
			return
		}
	}

	authRequiredGroup := apiEngine.Group("/")
	authRequiredGroup.Use(middlewares.Authorization)
	for _, endpoint := range needAuthorizationEndpoints {
		switch endpoint.Method {
		case http.MethodGet:
			authRequiredGroup.GET(endpoint.Path, endpoint.HandlerFunc)
		case http.MethodPut:
			authRequiredGroup.PUT(endpoint.Path, endpoint.HandlerFunc)
		case http.MethodPost:
			authRequiredGroup.POST(endpoint.Path, endpoint.HandlerFunc)
		case http.MethodDelete:
			authRequiredGroup.DELETE(endpoint.Path, endpoint.HandlerFunc)
		default:
			err = fmt.Errorf("Endpoint %s uses an unknown method %s", endpoint.Path, endpoint.Method)
			return
		}
	}

	return
}
