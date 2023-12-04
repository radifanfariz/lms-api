package models

import "time"

type ModuleMetadata struct {
	ID              int       `gorm:"primaryKey;column:n_id"`
	Slug            string    `gorm:"column:c_slug"`
	Name            string    `gorm:"column:c_name"`
	Description     string    `gorm:"column:c_description"`
	Src             string    `gorm:"column:c_src"`
	LearningJourney string    `gorm:"column:c_learning_journey"`
	Category        string    `gorm:"column:c_category"`
	MaxMonth        int       `gorm:"column:n_max_month"`
	CreatedBy       string    `gorm:"column:c_created_by"`
	UpdatedBy       string    `gorm:"column:c_updated_by"`
	CreatedAt       time.Time `gorm:"column:d_created_at"`
	UpdatedAt       time.Time `gorm:"column:d_updated_at"`
}

type ModuleData struct {
	ID             int    `gorm:"primaryKey;column:n_id"`
	Slug           string `gorm:"column:c_slug"`
	ModuleMetaID   int    `gorm:"column:n_module_meta_id"`
	PretestMetaID  int    `gorm:"column:n_pretest_meta_id"`
	PretestID      int    `gorm:"column:n_pretest_id"`
	MateriMetaID   int    `gorm:"column:n_materi_meta_id"`
	MateriID       int    `gorm:"column:n_materi_id"`
	PosttestMetaID int    `gorm:"column:n_posttest_meta_id"`
	PosttestID     int    `gorm:"column:n_posttest_id"`
	UserID         int    `gorm:"column:n_user_id"`
	GradeID        int    `gorm:"column:n_grade_id"`
}
