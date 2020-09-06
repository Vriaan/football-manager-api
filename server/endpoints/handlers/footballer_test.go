package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github/vriaan/footballmanagerapi/models"
	"github/vriaan/footballmanagerapi/tests"
)

func TestGetFootballerNotFound(t *testing.T) {
	nonExistentFootballerID := "20000000"
	getFootballerRoute := "/footballers/" + nonExistentFootballerID
	params := tests.TestParams{
		PathParams: gin.Params{
			gin.Param{Key: "id", Value: nonExistentFootballerID},
		},
	}
	responseStatus, responseBody, err := tests.TestHTTPHandler("GET", getFootballerRoute,
		params, GetFootballer)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusNotFound, responseStatus, string(responseBody))
	assert.Equal(t, fmt.Sprintf("{\"error\":\"%s\"}", http.StatusText(http.StatusNotFound)), string(responseBody))
}

func TestGetFootballer(t *testing.T) {
	footballerID := "1"
	getFootballerRoute := "/footballers/" + footballerID
	params := tests.TestParams{
		PathParams: gin.Params{
			gin.Param{Key: "id", Value: footballerID},
		},
	}
	responseStatus, responseBody, err := tests.TestHTTPHandler("GET", getFootballerRoute,
		params, GetFootballer)
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

	assert.Equal(t, http.StatusOK, responseStatus, string(responseBody))

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
	params := tests.TestParams{}
	responseStatus, responseBody, err := tests.TestHTTPHandler("GET", listFootballersRoute, params, ListFootballers)
	if err != nil {
		t.Fatal(err)
	}

	expectedFootballers := models.Footballers{}
	// total number of footballer matching
	expectedCount := models.GetDB().Find(&expectedFootballers).RowsAffected

	expectedResponse := struct {
		Count int64              `json:"count"`
		List  models.Footballers `json:"list"`
	}{Count: expectedCount, List: expectedFootballers}

	expectedJSONResponse, err := json.Marshal(&expectedResponse)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, responseStatus, string(responseBody))
	assert.Equal(t, string(expectedJSONResponse), string(responseBody))
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

	assert.Equal(t, http.StatusCreated, responseStatus, string(responseBody))
	assert.Equal(t, bodyParams["FirstName"], newFootballer.FirstName)
	assert.Equal(t, bodyParams["LastName"], newFootballer.LastName)
}

func TestDeleteFootballer(t *testing.T) {
	var err error
	newFootballerToDeleteAfter := models.Footballer{FirstName: "Test", LastName: "Delete"}
	if err = models.GetDB().Create(&newFootballerToDeleteAfter).Error; err != nil {
		t.Fatal(err)
	}
	toDeleteFootballerID := strconv.FormatUint(uint64(newFootballerToDeleteAfter.ID), 10)
	deleteFootballerPath := "/footballers/" + toDeleteFootballerID
	params := tests.TestParams{
		PathParams: gin.Params{
			gin.Param{Key: "id", Value: toDeleteFootballerID},
		},
	}
	responseStatus, responseBody, err := tests.TestHTTPHandler("DELETE", deleteFootballerPath, params, DeleteFootballer)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusNoContent, responseStatus, string(responseBody))
	assert.Equal(t, "", string(responseBody))

	err = models.GetDB().First(&models.Footballer{}, toDeleteFootballerID).Error
	assert.Error(t, gorm.ErrRecordNotFound, err)
}

func TestUpdateFootballer(t *testing.T) {
	var err error
	newFootballerToUpdateAfter := models.Footballer{FirstName: "Test", LastName: "ToUpdate"}
	if err = models.GetDB().Create(&newFootballerToUpdateAfter).Error; err != nil {
		t.Fatal(err)
	}
	toUpdateFootballerID := strconv.FormatUint(uint64(newFootballerToUpdateAfter.ID), 10)
	updateFootballerPath := "/footballers/" + toUpdateFootballerID
	params := tests.TestParams{
		PathParams: gin.Params{
			gin.Param{Key: "id", Value: toUpdateFootballerID},
		},
		QueryParams: gin.Params{
			gin.Param{Key: "LastName", Value: "Updated"},
		},
	}
	responseStatus, responseBody, err := tests.TestHTTPHandler("PUT", updateFootballerPath, params, UpdateFootballer)
	if err != nil {
		t.Fatal(err)
	}

	updatedFootballer := models.Footballer{}
	err = models.GetDB().First(&updatedFootballer, toUpdateFootballerID).Error
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, responseStatus, string(responseBody))
	assert.Equal(t, "Updated", updatedFootballer.LastName)
	var expectedResponse []byte
	if expectedResponse, err = json.Marshal(&updatedFootballer); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, string(expectedResponse), string(responseBody))
}

func BenchmarkGetFootballer(b *testing.B) {
	b.ReportAllocs()
	footballerID := "1"
	getFootballerPath := "/footballers/" + footballerID
	params := tests.TestParams{
		PathParams: gin.Params{
			gin.Param{Key: "id", Value: footballerID},
		},
	}
	for n := 0; n < b.N; n++ {
		_, _, err := tests.TestHTTPHandler("GET", getFootballerPath,
			params, GetFootballer)
		if err != nil {
			b.Fatal(err)
		}
	}
}
