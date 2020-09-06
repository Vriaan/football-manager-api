package models

import (
	"github.com/jinzhu/gorm"
)

// Manager represents model for table managers
type Manager struct {
	gorm.Model
	TeamID    uint
	FirstName string
	LastName  string
	Password  string
	Email     string
}
