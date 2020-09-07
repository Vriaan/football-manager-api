package actions

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github/vriaan/footballmanagerapi/models"
	"github/vriaan/footballmanagerapi/test"
)

func TestLogin(t *testing.T) {
	manager := models.Manager{}
	errorDB := models.GetDB().First(&manager).Error
	if errorDB != nil {
		t.Fatal(errorDB)
	}
	loginRoute := "/login"
	params := test.Params{
		BodyParams: map[string]interface{}{
			"Email":    manager.Email,
			"Password": manager.Password,
		},
	}
	responseStatus, responseBody, err := test.CallAction("POST", loginRoute,
		params, Login)
	if err != nil {
		t.Fatal(err)
	}
	response := struct {
		Token string         `json:"token"`
		User  models.Manager `json:"user"`
	}{}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, responseStatus, string(responseBody))
	assert.NotEmpty(t, response.Token)
	assert.Equal(t, manager.ID, response.User.ID)
	assert.Equal(t, manager.TeamID, response.User.TeamID)
	assert.Equal(t, manager.Email, response.User.Email)
	assert.Equal(t, manager.FirstName, response.User.FirstName)
	assert.Equal(t, manager.LastName, response.User.LastName)
}
