package serverapi

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github/vriaan/footballmanagerapi/endpoints"
	"github/vriaan/footballmanagerapi/endpoints/actions"
	"github/vriaan/footballmanagerapi/middlewares"
	"github/vriaan/footballmanagerapi/test"
)

func TestInitializeServer(t *testing.T) {
	db := test.GetDBConnection()
	serverAPI, err := Initialize(nil, db, ":8080", "/tmp/api.log")
	assert.Nil(t, err)
	assert.NotNil(t, serverAPI)
	assert.Len(t, serverAPI.Engine.Routes(), 0)

	apiEndpoints := endpoints.Endpoints{
		endpoints.Endpoint{http.MethodGet, "/toto",
			middlewares.Middlewares{},
			actions.Action(func(c *gin.Context) {}),
		},
		endpoints.Endpoint{http.MethodGet, "/tata",
			middlewares.Middlewares{func(c *gin.Context) {}, func(c *gin.Context) {}},
			actions.Action(func(c *gin.Context) {}),
		},
	}

	serverAPI, err = Initialize(&apiEndpoints, db, ":8080", "/tmp/api.log")
	assert.Nil(t, err)
	assert.NotNil(t, serverAPI)
	assert.Len(t, serverAPI.Engine.Routes(), len(apiEndpoints))

	_, err = Initialize(nil, nil, ":8080", "/tmp/api.log")
	assert.EqualError(t, err, "Nil database connection provided")
}
