package models

import "time"

type CertificateMasterData struct {
	ID        int       `gorm:"primaryKey;column:n_id" json:"id"`
	Name      string    `gorm:"column:c_name" json:"name"`
	Src       string    `gorm:"column:c_src" json:"src"`
	IsActive  *bool     `gorm:"default:false;column:b_isactive" json:"is_active"`
	CreatedBy string    `gorm:"column:c_created_by" json:"created_by"`
	UpdatedBy string    `gorm:"column:c_updated_by" json:"updated_by"`
	CreatedAt time.Time `gorm:"default:now();column:d_created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:now();column:d_updated_at" json:"updated_at"`
}

type CertificateUserData struct {
	ID         int        `gorm:"primaryKey;column:n_id" json:"id"`
	VerifiedID string     `gorm:"column:c_verified_id" json:"verified_id"`
	UserID     int        `gorm:"column:n_user_id" json:"user_id"`
	GlobalID   string     `gorm:"column:c_global_id" json:"global_id"`
	ModuleID   int        `gorm:"column:n_module_id" json:"module_id"`
	CreatedBy  string     `gorm:"column:c_created_by" json:"created_by"`
	UpdatedBy  string     `gorm:"column:c_updated_by" json:"updated_by"`
	CreatedAt  time.Time  `gorm:"default:now();column:d_created_at" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"default:now();column:d_updated_at" json:"updated_at"`
	UserData   UserData   `gorm:"foreignKey:UserID" json:"user_data"`
	ModuleData ModuleData `gorm:"foreignKey:GlobalID;references:GlobalID" json:"module_data"`
}
