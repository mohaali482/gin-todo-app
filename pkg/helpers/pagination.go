package helpers

import (
	"math"

	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int         `json:"limit,omitempty" form:"limit"`
	Page       int         `json:"page,omitempty" form:"page"`
	Count      int64       `json:"count"`
	TotalPages int         `json:"total_pages"`
	Items      interface{} `json:"items"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var count int64
	db.Model(value).Count(&count)
	pagination.Count = count
	totalPages := (math.Ceil(float64(count) / float64(pagination.Limit)))
	pagination.TotalPages = int(totalPages)
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}
