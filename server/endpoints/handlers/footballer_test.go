package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"

	"github/vriaan/footballmanagerapi/models"
	"github/vriaan/footballmanagerapi/tests"
)

func TestGetFootballerNotFound(t *testing.T) {
	nonExistentFootballerID := "20000000"
	getFootballerRoute := "/footballers/" + nonExistentFootballerID
	contextParameters := tests.TestParams{
		PathParams: gin.Params{
			gin.Param{Key: "id", Value: nonExistentFootballerID},
		},
	}
	responseStatus, responseBody, err := tests.TestHTTPHandler("GET", getFootballerRoute,
		contextParameters, GetFootballer)
	if err != nil {
		t.Fatal(err)
	}

	description := fmt.Sprintf("Status code %d from call %s", http.StatusNotFound, getFootballerRoute)
	assert.Equal(t, http.StatusNotFound, responseStatus, description)
	assert.Equal(t, fmt.Sprintf("{\"error\":\"%s\"}", http.StatusText(http.StatusNotFound)), string(responseBody))
}

func TestGetFootballer(t *testing.T) {
	footballerID := "1"
	getFootballerRoute := "/footballers/" + footballerID
	contextParameters := tests.TestParams{
		PathParams: gin.Params{
			gin.Param{Key: "id", Value: footballerID},
		},
	}
	responseStatus, responseBody, err := tests.TestHTTPHandler("GET", getFootballerRoute,
		contextParameters, GetFootballer)
	if err != nil {
		t.Fatal(err)
	}

	footballerFound := models.Footballer{}
	if err := json.Unmarshal(responseBody, &footballerFound); err != nil {
		t.Fatal(err)
	}
	if (models.Footballer{}) == footballerFound {
		t.Fatalf("Unmarshalling JSON could not match response to expected data type, response was: %s", responseBody)
	}

	description := fmt.Sprintf("Status code %d from call %s", http.StatusOK, getFootballerRoute)
	assert.Equal(t, http.StatusOK, responseStatus, description)

	expectedFootballerData := models.Footballer{}
	err = models.GetDB().First(&expectedFootballerData, footballerID).Error
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedFootballerData.ID, footballerFound.ID)
	assert.Equal(t, expectedFootballerData.FirstName, footballerFound.FirstName)
	assert.Equal(t, expectedFootballerData.LastName, footballerFound.LastName)
}

func TestListFootballer(t *testing.T) {
	listFootballersRoute := "/footballers"
	contextParameters := tests.TestParams{}
	responseStatus, responseBody, err := tests.TestHTTPHandler("GET", listFootballersRoute, contextParameters, ListFootballers)
	if err != nil {
		t.Fatal(err)
	}

	response := struct {
		Count int64              `json:"count"`
		List  models.Footballers `json:"list"`
	}{}

	if err := json.Unmarshal(responseBody, &response); err != nil {
		t.Fatal(err)
	}

	expectedFootballers := models.Footballers{}
	// total number of footballer matching
	expectedCount := models.GetDB().Find(&expectedFootballers).RowsAffected
	description := fmt.Sprintf("Status code %d from call %s", http.StatusOK, listFootballersRoute)
	assert.Equal(t, http.StatusOK, responseStatus, description)
	assert.Equal(t, expectedCount, response.Count)

	assert.True(t, expectedFootballers == response.List)
}

func TestRegisterNewFootballer(t *testing.T) {
	createFootballerPath := "/footballers"

	bodyParams := make(map[string]interface{}, 0)
	bodyParams["FirstName"] = "Blade"
	bodyParams["LastName"] = "Runner"
	contextParams := tests.TestParams{BodyParams: bodyParams}
	responseStatus, responseBody, err := tests.TestHTTPHandler("POST", createFootballerPath, contextParams, RegisterNewFootballer)
	if err != nil {
		t.Fatal(err)
	}

	newFootballer := models.Footballer{}
	if err := json.Unmarshal(responseBody, &newFootballer); err != nil {
		t.Fatal(err)
	}
	if (models.Footballer{}) == newFootballer {
		t.Fatalf("Unmarshalling JSON could not match response to expected data type, response was: %s", responseBody)
	}

	assert.Equal(t,
		http.StatusOK, responseStatus,
		fmt.Sprintf("Status code %d from call %s", http.StatusOK, createFootballerPath),
	)
	assert.Equal(t, bodyParams["FirstName"], newFootballer.FirstName)
	assert.Equal(t, bodyParams["LastName"], newFootballer.LastName)
}

func BenchmarkGetFootballer(b *testing.B) {
	b.ReportAllocs()
	footballerID := "1"
	getFootballerPath := "/footballers/" + footballerID
	contextParameters := tests.TestParams{
		PathParams: gin.Params{
			gin.Param{Key: "id", Value: footballerID},
		},
	}
	for n := 0; n < b.N; n++ {
		_, _, err := tests.TestHTTPHandler("GET", getFootballerPath,
			contextParameters, GetFootballer)
		if err != nil {
			b.Fatal(err)
		}
	}
}
