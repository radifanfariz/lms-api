package models

import (
	"time"
)

type GalleryData struct {
	ID        int       `gorm:"primaryKey;column:n_id" json:"id"`
	UserID    int       `gorm:"column:n_user_id" json:"user_id"`
	Name      string    `gorm:"column:c_name" json:"name"`
	Src       string    `gorm:"column:c_src" json:"src"`
	CreatedBy string    `gorm:"column:c_created_by" json:"created_by"`
	UpdatedBy string    `gorm:"column:c_updated_by" json:"updated_by"`
	CreatedAt time.Time `gorm:"default:now();column:d_created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:now();column:d_updated_at" json:"updated_at"`
	UserData  UserData  `gorm:"foreignKey:UserID" json:"user_data"`
}
