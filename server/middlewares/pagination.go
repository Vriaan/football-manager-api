package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const (
	// PaginationSizeKey is page result size query parameter name
	PaginationSizeKey = "limit"
	// DefaultPaginationSize is page result size per default
	DefaultPaginationSize = "100"
	// PaginationOffsetKey is page result offset query parameter name
	PaginationOffsetKey = "offset"
	// DefaultOffset is page result offset per default
	DefaultOffset = "0"
)

// Paginate gets a database handler and adds the limit & offset from asked from request or set default
func Paginate(c *gin.Context, db *gorm.DB) *gorm.DB {
	limit := c.DefaultQuery(PaginationSizeKey, DefaultPaginationSize)
	offset := c.DefaultQuery(PaginationOffsetKey, DefaultOffset)
	// if pageSize, err = strconv.ParseUint(c.DefaultQuery(PaginationSizeKey, DefaultPaginationSize), 10, 64); err != nil {
	// 	pageSize = DefaultPaginationSize
	// }
	// if pageOffset, err = strconv.ParseUint(c.DefaultQuery(ResultOffsetKey, DefaultOffset), 10, 64); err != nil {
	// 	pageOffset = 0
	// }
	return db.Offset(offset).Limit(limit)
	// return func(db *gorm.DB) *gorm.DB {
	// 	return db.Offset(offset).Limit(limit)
	// }
}
