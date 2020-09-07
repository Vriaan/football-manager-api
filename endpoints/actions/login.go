package actions

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github/vriaan/footballmanagerapi/middlewares"
	"github/vriaan/footballmanagerapi/models"
)

// Login route checks user credentials and returns a validation token
func Login(c *gin.Context) {
	var (
		searchedManager, foundManager models.Manager
		authorizationToken            string
		err                           error
	)

	if err = c.ShouldBindJSON(&searchedManager); err != nil {
		abortError(c, http.StatusBadRequest, err)
		return
	}

	foundManager, err = searchedManager.First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			abortStatus(c, http.StatusForbidden)
		} else {
			abortError(c, http.StatusInternalServerError, err)
		}
		return
	}

	// Generate an authentication token and attach it to the response
	// err = middlewares.AttachAuthenticationToken(c, manager.ID)
	authorizationToken, err = middlewares.CreateAuthToken(foundManager.ID)
	if err != nil {
		abortError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": authorizationToken,
		"user": gin.H{
			"ID":        foundManager.ID,
			"TeamID":    foundManager.TeamID,
			"Email":     foundManager.Email,
			"FirstName": foundManager.FirstName,
			"LastName":  foundManager.LastName,
		},
	})
}
