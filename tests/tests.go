package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// Initialize MySQL driver for tests
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"

	"github/vriaan/footballmanagerapi/models"
	"github/vriaan/footballmanagerapi/server/endpoints"
)

// Creates some singleton used later for the tests
func init() {
	GetDbConnection()
	GetAPIEngine()
}

var (
	testAPI   *gin.Engine
	dbHandler *gorm.DB
)

// GetAPIEngine provides a test api engine within a singleton
func GetAPIEngine() *gin.Engine {
	if testAPI != nil {
		return testAPI
	}
	_, testAPI = gin.CreateTestContext(httptest.NewRecorder())
	err := endpoints.Register(testAPI)
	if err != nil {
		panic(err)
	}

	return testAPI
}

// GetDbConnection instanciates a database connection for test within a singleton (TODO: Use a SQL Mock ?)
func GetDbConnection() *gorm.DB {
	if dbHandler != nil {
		return dbHandler
	}

	var err error
	dbHandler, err = gorm.Open(
		"mysql",
		"root:root@tcp(football-manager-db-test)/footballmanager_test?charset=utf8&parseTime=True&loc=Local",
	)
	if err != nil {
		panic(err)
	}
	models.SetDb(dbHandler)

	return dbHandler
}

// DataToBufferizedJSON handles data transformation to JSON and bufferize it
func dataToBufferizedJSON(data interface{}) (dataBuffized io.Reader, err error) {
	if data == nil {
		return
	}

	var jsonData []byte
	jsonData, err = json.Marshal(data)
	if err != nil {
		err = errors.Wrap(err, "Error Marshalling test data structure to json")
		return
	}
	if len(jsonData) == 0 {
		err = errors.New("Error Marshalling test data structure to json: Returned empty data")
		return
	}
	dataBuffized = bytes.NewBuffer(jsonData)
	return
}

// DoJSONRequest creates and performs a JSON request to the API router
func DoJSONRequest(method, url string, parameters interface{}) (responseStatus int, responseBody []byte, err error) {
	var (
		request *http.Request
		data    io.Reader
	)

	if data, err = dataToBufferizedJSON(parameters); err != nil {
		return
	}
	if request, err = http.NewRequest(method, url, data); err != nil {
		return
	}
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	testAPI.ServeHTTP(recorder, request)
	result := recorder.Result()

	responseStatus = result.StatusCode
	responseBody, err = ioutil.ReadAll(result.Body)
	return
}
