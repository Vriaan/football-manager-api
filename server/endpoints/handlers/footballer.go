package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github/vriaan/footballmanagerapi/models"
)

// RegisterNewFootballer registers a new footballer to database
func RegisterNewFootballer(c *gin.Context) {
	var err error
	var newFootballer models.Footballer

	if err = c.ShouldBindJSON(&newFootballer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = models.GetDb().Create(&newFootballer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newFootballer)
}

// GetFootballers returns footballers matching provided criterias
func GetFootballers(c *gin.Context) {
}

// GetFootballer gets a new footballer to database
func GetFootballer(c *gin.Context) {
	footballerID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	footballer := &models.Footballer{}
	err = models.GetDb().First(footballer, footballerID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": http.StatusText(http.StatusNotFound)})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, footballer)
}
