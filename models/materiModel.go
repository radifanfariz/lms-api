package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type MateriMetadata struct {
	ID          int       `gorm:"primaryKey;column:n_id" json:"id"`
	ModuleID    int       `gorm:"column:n_module_id" json:"module_id"`
	GlobalID    string    `gorm:"column:c_global_id" json:"global_id"`
	Name        string    `gorm:"column:c_name" json:"name"`
	Description string    `gorm:"column:c_description" json:"description"`
	Src         string    `gorm:"column:c_src" json:"src"`
	CreatedBy   string    `gorm:"column:c_created_by" json:"created_by"`
	UpdatedBy   string    `gorm:"column:c_updated_by" json:"updated_by"`
	CreatedAt   time.Time `gorm:"default:now();column:d_created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:now();column:d_updated_at" json:"updated_at"`
}

type MateriData struct {
	ID           int            `gorm:"primaryKey;column:n_id" json:"id"`
	ModuleID     int            `gorm:"column:n_module_id" json:"module_id"`
	MateriMetaID int            `gorm:"column:n_materi_meta_id" json:"materi_meta_id"`
	GlobalID     string         `gorm:"column:c_global_id" json:"global_id"`
	Name         string         `gorm:"column:c_name" json:"name"`
	Description  string         `gorm:"column:c_description" json:"description"`
	Type         string         `gorm:"column:c_type" json:"type"`
	Src          string         `gorm:"column:c_src" json:"src"`
	IsPublished  *bool          `gorm:"column:b_ispublished" json:"is_published"`
	CreatedBy    string         `gorm:"column:c_created_by" json:"created_by"`
	UpdatedBy    string         `gorm:"column:c_updated_by" json:"updated_by"`
	CreatedAt    time.Time      `gorm:"default:now();column:d_created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"default:now();column:d_updated_at" json:"updated_at"`
	Metadata     MateriMetadata `gorm:"foreignKey:MateriMetaID" json:"metadata"`
}
type MateriResultData struct {
	ID         int              `gorm:"primaryKey;column:n_id" json:"id"`
	UserID     int              `gorm:"column:n_user_id" json:"user_id"`
	MateriID   int              `gorm:"column:n_materi_id" json:"materi_id"`
	GlobalID   string           `gorm:"column:c_global_id" json:"global_id"`
	Start      pgtype.Timestamp `gorm:"column:d_start" json:"start"`
	End        pgtype.Timestamp `gorm:"column:d_end" json:"end"`
	Duration   Duration         `gorm:"column:d_duration" json:"duration"`
	CreatedBy  string           `gorm:"column:c_created_by" json:"created_by"`
	UpdatedBy  string           `gorm:"column:c_updated_by" json:"updated_by"`
	CreatedAt  time.Time        `gorm:"default:now();column:d_created_at" json:"created_at"`
	UpdatedAt  time.Time        `gorm:"default:now();column:d_updated_at" json:"updated_at"`
	UserData   UserData         `gorm:"foreignKey:UserID" json:"user_data"`
	MateriData MateriData       `gorm:"foreignKey:MateriID" json:"materi_data"`
	ModuleData ModuleData       `gorm:"foreignKey:GlobalID;references:GlobalID" json:"module_data"`
}
