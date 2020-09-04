package server

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github/vriaan/footballmanagerapi/tests"
)

func TestInitializeServer(t *testing.T) {
	db := tests.GetDbConnection()
	serverAPI, err := Initialize(nil, db, ":8080", "/tmp/api.log")
	assert.Nil(t, err)
	assert.NotNil(t, serverAPI)

	serverAPI, err = Initialize(gin.New(), db, ":8080", "/tmp/api.log")
	assert.Nil(t, err)
	assert.NotNil(t, serverAPI)
}
