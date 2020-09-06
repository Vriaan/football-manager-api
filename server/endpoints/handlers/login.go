package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github/vriaan/footballmanagerapi/models"
	"github/vriaan/footballmanagerapi/server/middlewares"
)

// Login route checks user credentials and returns a token validation
func Login(c *gin.Context) {
	var err error
	var manager models.Manager

	if err = c.ShouldBindJSON(&manager); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.GetDB().Where(&manager).First(&manager).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": http.StatusText(http.StatusNotFound)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Generate an authentication token and attach it to the response
	err = middlewares.AttachAuthenticationToken(c, manager.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, manager)
}
