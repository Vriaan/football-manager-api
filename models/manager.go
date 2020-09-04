package models

import (
	"github.com/jinzhu/gorm"
)

// Manager represents model for table managers
type Manager struct {
	gorm.Model
	Name  string `json:"name"`
	Level uint64 `json:"level"`
}
