package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github/vriaan/footballmanagerapi/models"
	"github/vriaan/footballmanagerapi/server/middlewares"
)

// Login route checks user credentials and returns a validation token
func Login(c *gin.Context) {
	var (
		managerSearch, foundManager models.Manager
		authorizationToken          string
		err                         error
	)

	if err = c.ShouldBindJSON(&managerSearch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.GetDB().Where(&managerSearch).First(&foundManager).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusForbidden, gin.H{"error": http.StatusText(http.StatusForbidden)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Generate an authentication token and attach it to the response
	// err = middlewares.AttachAuthenticationToken(c, manager.ID)
	authorizationToken, err = middlewares.CreateAuthToken(foundManager.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
