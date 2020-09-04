package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github/vriaan/footballmanagerapi/models"
	"github/vriaan/footballmanagerapi/tests"
)

func TestGetFootballer(t *testing.T) {
	getFootballerRoute := "/footballers/2000000000000"
	responseStatus, responseBody, err := tests.DoJSONRequest("GET", getFootballerRoute, nil)
	if err != nil {
		t.Fatal(err)
	}

	description := fmt.Sprintf("Status code %d from call %s", http.StatusNotFound, getFootballerRoute)
	assert.Equal(t, http.StatusNotFound, responseStatus, description)
	assert.Equal(t, fmt.Sprintf("{\"error\":\"%s\"}", http.StatusText(http.StatusNotFound)), string(responseBody))

}

func TestPostFootballer(t *testing.T) {
	createFootballerPath := "/footballers"
	params := struct {
		FirstName string
		LastName  string
		TeamID    uint64
	}{"Blade", "Runner", 1}

	responseStatus, responseBody, err := tests.DoJSONRequest("POST", createFootballerPath, params)
	if err != nil {
		t.Fatal(err)
	}

	newFootballer := models.Footballer{}
	if err := json.Unmarshal(responseBody, &newFootballer); err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, models.Footballer{}, newFootballer)

	assert.Equal(t,
		http.StatusOK, responseStatus,
		fmt.Sprintf("Status code %d from call %s", http.StatusOK, createFootballerPath),
	)

	assert.Equal(t, params.FirstName, newFootballer.FirstName)
	assert.Equal(t, params.LastName, newFootballer.LastName)
}

func BenchmarkGetFootballer(b *testing.B) {
	b.ReportAllocs()
	getFootballerPath := "/footballers/1"
	for n := 0; n < b.N; n++ {
		_, _, err := tests.DoJSONRequest("GET", getFootballerPath, nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}
