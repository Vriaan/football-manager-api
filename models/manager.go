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

// First retruns the first record maching condition on non empty field from m
func (m *Manager) First() (foundManager Manager, err error) {
	err = GetDB().Where(m).First(&foundManager).Error
	return
}
