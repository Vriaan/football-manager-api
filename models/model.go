package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
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

// InitDatabaseConnection sets models database connection pool singleton and set it up
func InitDatabaseConnection(dbType, dsn string) (err error) {
	// Just ignore
	if databaseConnectionPoolHandler != nil {
		return
	}

	if databaseConnectionPoolHandler, err = gorm.Open(dbType, dsn); err != nil {
		err = errors.Wrapf(err, "Initialize Database connection pool handler")
	}
	// TODO Add some settings about pool (pool size, connection expiry, etc)
	return
}
