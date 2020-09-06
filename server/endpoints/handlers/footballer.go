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

	if err = models.GetDB().Create(&newFootballer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newFootballer)
}

// DeleteFootballer removes footballer from database (no soft delete here)
func DeleteFootballer(c *gin.Context) {

}

// ListFootballers returns footballers matching provided criterias
func ListFootballers(c *gin.Context) {
	var (
		footballer       models.Footballer
		footballers      models.Footballers
		footballersCount int
		err              error
	)
	if err = c.ShouldBindQuery(&footballer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// err = middlewares.Paginate(c, models.GetDB()).Find(&footballers, searchCondition).Error
	footballers, err = footballer.Find()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": http.StatusText(http.StatusNotFound)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	footballersCount, err = footballer.Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"list":  footballers,
		"count": footballersCount,
	})
}

// GetFootballer gets a new footballer to database
func GetFootballer(c *gin.Context) {
	footballerID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	footballer := &models.Footballer{}
	err = models.GetDB().First(footballer, footballerID).Error
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
