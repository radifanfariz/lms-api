package utils

import (
	"gorm.io/gorm"
)

type Pagination struct {
	Limit        int         `json:"per_page,omitempty" query:"per_page"`
	Page         int         `json:"page,omitempty" query:"page"`
	Sort         string      `json:"sort,omitempty" query:"sort"`
	FilterColumn string      `json:"filter_column,omitempty" query:"filter_column"`
	Filter       string      `json:"filter,omitempty" query:"filter"`
	TotalData    int64       `json:"total_data"`
	TotalPages   int         `json:"total_pages"`
	Data         interface{} `json:"data"`
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

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "n_id desc"
	}
	return p.Sort
}

func (p *Pagination) GetFilter() string {
	if p.Filter == "" {
		p.Filter = ""
	}
	return p.Filter
}

func (p *Pagination) GetFilterColumn() string {
	if p.FilterColumn == "" {
		p.FilterColumn = ""
	}
	return p.FilterColumn
}

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if pagination.GetFilter() != "" && pagination.GetFilterColumn() != "" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Where(pagination.GetFilterColumn(), pagination.GetFilter())
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
