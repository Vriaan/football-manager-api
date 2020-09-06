package endpoints

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	endpoints := append(freeAccessEndpoints, needAuthorizationEndpoints...)
	endpointsDeclaredCopy := Get()
	assert.NotNil(t, endpointsDeclaredCopy)
	assert.NotEqual(t, endpointsDeclaredCopy, endpoints, "Get returns a copy of the endpoints declared")
	assert.Len(t, endpointsDeclaredCopy, len(freeAccessEndpoints)+len(needAuthorizationEndpoints))

	for index, route := range endpointsDeclaredCopy {
		sameRoute := endpoints[index].Method == route.Method &&
			endpoints[index].Path == route.Path &&
			endpoints[index].Handler == route.Handler &&
			route.Handler == "" &&
			route.HandlerFunc != nil
		assert.True(t, sameRoute)
	}
}

func TestRegister(t *testing.T) {
	endpoints := Get()
	engine := gin.New()
	Register(engine)
	engineRouteRegistered := engine.Routes()
	assert.Len(t, engineRouteRegistered, len(endpoints))
	// for index, route := range engineRouteRegistered {
	// 	sameRouteRegistered := endpoints[index].Method == route.Method &&
	// 		endpoints[index].Path == route.Path &&
	// 		endpoints[index].Handler != route.Handler &&
	// 		route.HandlerFunc != nil
	//
	// 	assert.True(t, sameRouteRegistered)
	// }

}
