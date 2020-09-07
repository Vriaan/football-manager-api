package test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	// Initialize MySQL driver for tests
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github/vriaan/footballmanagerapi/models"
)

// TODO Replace by a mock SQL
var dbHandler *gorm.DB

// Params is a convenient structure to pass parameters for tests endpoint
type Params struct {
	PathParams  gin.Params
	QueryParams gin.Params
	BodyParams  map[string]interface{}
}

// CallAction wraps logic to test http handler response (a bit too complicated for just testing .... (wait for gin to make it easy))
func CallAction(
	method, urlEndpoint string,
	params Params,
	apiHandler gin.HandlerFunc,
) (responseStatus int, responseBody []byte, err error) {
	GetDBConnection()
	recorder := httptest.NewRecorder()
	handler := func(writer http.ResponseWriter, request *http.Request) {

		gin.SetMode(gin.TestMode)
		context, _ := gin.CreateTestContext(writer)
		context.Request = request
		context.Params = append(params.PathParams, params.QueryParams...)
		apiHandler(context)
	}

	// Need to rework the parameters for request data to be Marshalled as JSON
	requestData := make(map[string]interface{}, 0)
	for _, param := range params.QueryParams {
		requestData[param.Key] = param.Value
	}
	for key, value := range params.BodyParams {
		requestData[key] = value
	}
	var data io.Reader
	if data, err = dataToBufferizedJSON(requestData); err != nil {
		return
	}

	request := httptest.NewRequest(method, urlEndpoint, data)
	request.Header.Add("Content-Type", "application/json")

	handler(recorder, request)
	result := recorder.Result()
	responseStatus = result.StatusCode
	responseBody, err = ioutil.ReadAll(result.Body)
	return
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

// GetDBConnection instanciates a database connection for test within a singleton (TODO: Use a SQL Mock ?)
func GetDBConnection() *gorm.DB {
	if dbHandler != nil {
		return dbHandler
	}

	var err error
	dbHandler, err = gorm.Open("mysql", os.Getenv("DB_DSN"))
	if err != nil {
		panic("Open test db handler error: " + err.Error())
	}
	models.SetDb(dbHandler)

	return dbHandler
}
