package actions

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
		abortError(c, http.StatusBadRequest, err)
		return
	}

	if err = models.GetDB().Create(&newFootballer).Error; err != nil {
		abortError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, newFootballer)
}

// DeleteFootballer removes footballer from database (no soft delete here)
func DeleteFootballer(c *gin.Context) {
	footballerID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		abortError(c, http.StatusBadRequest, err)
		return
	}

	deletedFootballer := &models.Footballer{}
	err = models.GetDB().Delete(&deletedFootballer, footballerID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			abortStatus(c, http.StatusNotFound)
		} else {
			abortError(c, http.StatusInternalServerError, err)
		}
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// UpdateFootballer updates footballer saved informations
func UpdateFootballer(c *gin.Context) {
	footballerID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		abortError(c, http.StatusBadRequest, err)
		return
	}

	var footballerFieldToUpdate models.Footballer
	if err = c.ShouldBindJSON(&footballerFieldToUpdate); err != nil {
		abortError(c, http.StatusBadRequest, err)
		return
	}

	// footballerFieldToUpdate.ID = 0
	var updatedFootballer models.Footballer
	updatedFootballer, err = (&models.Footballer{}).UpdateOne(uint(footballerID), footballerFieldToUpdate)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			abortStatus(c, http.StatusNotFound)
		} else {
			abortError(c, http.StatusInternalServerError, err)
		}
		return
	}

	c.JSON(http.StatusOK, updatedFootballer)
}

// ListFootballers returns footballers matching provided criterias
func ListFootballers(c *gin.Context) {
	var (
		FootballerConditions models.Footballer
		footballers          models.Footballers
		footballersCount     int
		err                  error
	)
	if err = c.ShouldBindQuery(&FootballerConditions); err != nil {
		abortError(c, http.StatusBadRequest, err)
		return
	}
	limit, offset := paginate(c)
	footballers, err = FootballerConditions.Find(limit, offset)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			abortStatus(c, http.StatusNotFound)
		} else {
			abortError(c, http.StatusInternalServerError, err)
		}
		return
	}
	footballersCount, err = FootballerConditions.Count()
	if err != nil {
		abortError(c, http.StatusInternalServerError, err)
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
		abortError(c, http.StatusBadRequest, err)
		return
	}

	footballer := &models.Footballer{}
	err = models.GetDB().First(footballer, footballerID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			abortStatus(c, http.StatusNotFound)
		} else {
			abortError(c, http.StatusInternalServerError, err)
		}
		return
	}

	c.JSON(http.StatusOK, footballer)
}
