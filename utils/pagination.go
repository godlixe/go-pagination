package utils

import (
	"fmt"
	"go-pagination/entity"
	"math"

	"gorm.io/gorm"
)

func PaginateOffset(totalRows int64, pagination *entity.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Order(pagination.GetSort()).Limit(pagination.GetLimit())
	}
}

func PaginateSeek(totalRows int64, pagination *entity.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	seekID := (pagination.Limit * (pagination.Page - 1)) - int(pagination.TotalRows)
	fmt.Println(seekID)
	if seekID > 0 {
		seekID = 0
	} else {
		seekID = int(math.Abs(float64(seekID)))
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(("id <= ?"), (seekID)).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
