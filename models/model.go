package models

import (
	"github.com/jinzhu/gorm"
)

var (
	databaseConnectionPoolHandler *gorm.DB
)

// // DataModel is supposed to be an interface of models to be able to generalize some CRUD query
// type DataModel interface {
// 	Name() string
// }

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

// // Find retrieves footballers record matching fields conditions setted up within f
// func Find(results *gorm.Model, conditions *gorm.Model, limit, offset string) (err error) {
// 	query := GetDB().Where(conditions)
// 	if limit == "" {
// 		query.Limit(limit)
// 	}
// 	if offset == "" {
// 		query.Offset(offset)
// 	}
// 	err = query.Find(results).Error
// 	return
// }
