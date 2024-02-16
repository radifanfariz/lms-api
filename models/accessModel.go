package models

import (
	"time"

	"github.com/lib/pq"
)

type AccessData struct {
	ID           int           `gorm:"primaryKey;column:n_id" json:"id"`
	ModuleMetaID int           `gorm:"column:n_module_meta_id" json:"module_meta_id"`
	GlobalID     string        `gorm:"column:c_global_id" json:"global_id"`
	ArrayGradeID pq.Int64Array `gorm:"column:n_array_grade_id;type:integer[]" json:"array_grade_id"`
	ArrayUserID  pq.Int64Array `gorm:"column:n_array_user_id;type:integer[]" json:"array_user_id"`
	CreatedBy    string        `gorm:"column:c_created_by" json:"created_by"`
	UpdatedBy    string        `gorm:"column:c_updated_by" json:"updated_by"`
	CreatedAt    time.Time     `gorm:"default:now();column:d_created_at" json:"created_at"`
	UpdatedAt    time.Time     `gorm:"default:now();column:d_updated_at" json:"updated_at"`
	ModuleData   ModuleData    `gorm:"foreignKey:ModuleMetaID" json:"module_metadata"`
}
