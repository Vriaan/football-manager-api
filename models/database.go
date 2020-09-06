package models

import (
	"github.com/jinzhu/gorm"
)

var (
	databaseConnectionPoolHandler *gorm.DB
)

// GetDB returns a MySQL database connection pool
func GetDB() *gorm.DB {
	if databaseConnectionPoolHandler == nil {
		panic("Database connection has not been instanciated")
	}
	return databaseConnectionPoolHandler
}

// SetDb sets models database connection pool singleton and set it up
func SetDb(dbHandler *gorm.DB) {
	databaseConnectionPoolHandler = dbHandler
	// TODO Add some settings about pool (pool size, connection expiry, etc)
}
