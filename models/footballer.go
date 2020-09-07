package models

import (
	"github.com/jinzhu/gorm"
)

// Footballers convenient type for Footballer slice
type Footballers []Footballer

// Footballer represents the model for table footballers
type Footballer struct {
	gorm.Model
	FirstName string
	LastName  string
}

// Find retrieves footballers record matching fields conditions setted up within f
func (f *Footballer) Find(limit, offset string) (footballers Footballers, err error) {
	query := GetDB().Where(f)
	if limit == "" {
		query.Limit(limit)
	}
	if offset == "" {
		query.Offset(offset)
	}
	err = query.Find(&footballers).Error
	return
}

// Count retrieves the records number matching conditions within f
func (f *Footballer) Count() (count int, err error) {
	err = GetDB().Model(f).Where(f).Count(&count).Error
	return
}

// UpdateOne updates the footballer informations matching the record ID provided
func (f *Footballer) UpdateOne(footballerID uint, fieldToUpdate Footballer) (footballer Footballer, err error) {
	if err = GetDB().Model(f).Where("id = ?", footballerID).Updates(fieldToUpdate).Error; err != nil {
		return
	}
	err = GetDB().First(&footballer, footballerID).Error
	return
}
