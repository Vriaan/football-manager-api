package models

import (
	"github.com/jinzhu/gorm"
)

// Manager represents model for table managers
type Manager struct {
	gorm.Model
	FirstName string
	LastName  uint64
	Password  string
}
