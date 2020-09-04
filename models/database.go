package models

import "github.com/jinzhu/gorm"

var (
	databaseConnectionPoolHandler *gorm.DB
)

// GetDb returns a MySQL database connection pool
func GetDb() *gorm.DB {
	if databaseConnectionPoolHandler == nil {
		panic("Database connection has not been instanciated")
	}
	return databaseConnectionPoolHandler
}

// SetDb sets models database connection pool singleton
func SetDb(dbHandler *gorm.DB) {
	databaseConnectionPoolHandler = dbHandler
}
