package models

import "time"

type AccessData struct {
	ID          int        `gorm:"primaryKey;column:n_id" json:"id"`
	ModuleID    int        `gorm:"column:n_module_id" json:"module_id"`
	GlobalID    string     `gorm:"column:n=c_global_id" json:"global_id"`
	JSONGradeID JSONB      `gorm:"type:jsonb;column:j_grade_id" json:"json_grade_id"`
	JSONUserID  JSONB      `gorm:"type:jsonb;column:j_user_id" json:"json_user_id"`
	CreatedBy   string     `gorm:"column:c_created_by" json:"created_by"`
	UpdatedBy   string     `gorm:"column:c_updated_by" json:"updated_by"`
	CreatedAt   time.Time  `gorm:"default:now();column:d_created_at" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:now();column:d_updated_at" json:"updated_at"`
	ModuleData  ModuleData `gorm:"foreignKey:ModuleID" json:"module_data"`
}
