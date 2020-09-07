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

// Create creates the current footballer with the structure non empty data set
func (f *Footballer) Create() (err error) {
	err = GetDB().Create(f).Error
	return
}

// List retrieves footballers record matching non empty fields setted up within f
func (f *Footballer) List(limit, offset string) (footballers Footballers, err error) {
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

// UpdateOne updates the footballer informations matching the record ID provided with the fields to update passed
func (f *Footballer) UpdateOne(footballerID uint, fieldToUpdate Footballer) (footballer Footballer, err error) {
	if err = GetDB().Model(f).Where("id = ?", footballerID).Updates(fieldToUpdate).Error; err != nil {
		return
	}
	err = GetDB().First(&footballer, footballerID).Error
	return
}

// Delete deletes the records matching information within f. If f is empty acts like a batch delete
func (f *Footballer) Delete() (err error) {
	err = GetDB().Delete(f).Error
	return
}
