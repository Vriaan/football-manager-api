package endpoints

import (
	"net/http"

	"github/vriaan/footballmanagerapi/endpoints/actions"
	"github/vriaan/footballmanagerapi/middlewares"
)

// Endpoints is a convenient type to means Endpoint slice
type Endpoints []Endpoint

// Endpoint represents an API endpoint
type Endpoint struct {
	// HTTP Method of the Restful API
	Method string
	// Endpoint URI
	Path string
	// List of middlewares to be called before the endpoint action
	Middlewares middlewares.Middlewares
	// Endpoint action
	Action actions.Action
}

// Get returns an endpoints to be registered within the API
func Get() Endpoints {
	return Endpoints{
		Endpoint{http.MethodGet, "/ping", middlewares.Middlewares{}, actions.Ping},
		Endpoint{http.MethodPost, "/login", middlewares.Middlewares{}, actions.Login},
		Endpoint{http.MethodPost, "/footballers", middlewares.Middlewares{middlewares.Authorization}, actions.RegisterNewFootballer},
		Endpoint{http.MethodGet, "/footballers", middlewares.Middlewares{middlewares.Authorization}, actions.ListFootballers},
		Endpoint{http.MethodGet, "/footballers/:id", middlewares.Middlewares{middlewares.Authorization}, actions.GetFootballer},
		Endpoint{http.MethodPut, "/footballers/:id", middlewares.Middlewares{middlewares.Authorization}, actions.UpdateFootballer},
		Endpoint{http.MethodDelete, "/footballers/:id", middlewares.Middlewares{middlewares.Authorization}, actions.DeleteFootballer},
	}
}
