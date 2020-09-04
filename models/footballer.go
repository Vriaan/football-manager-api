package models

import (
	"github.com/jinzhu/gorm"
)

// Footballer represents model for table footballers
type Footballer struct {
	gorm.Model
	TeamID    uint64 `binding:"required"`
	FirstName string `binding:"required"`
	LastName  string `binding:"required"`
}
