package models

import "time"

type CategoryData struct {
	ID        int       `gorm:"primaryKey;column:n_id" json:"id"`
	Domain    string    `gorm:"column:c_domain" json:"domain"`
	Label     string    `gorm:"column:c_label" json:"label"`
	Value     string    `gorm:"column:c_value" json:"value"`
	Seq       int       `gorm:"column:n_seq" json:"seq"`
	CreatedBy string    `gorm:"column:c_created_by" json:"created_by"`
	UpdatedBy string    `gorm:"column:c_updated_by" json:"updated_by"`
	CreatedAt time.Time `gorm:"default:now();column:d_created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:now();column:d_updated_at" json:"updated_at"`
}
