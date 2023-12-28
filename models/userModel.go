package models

import "time"

type UserData struct {
	ID           int    `gorm:"primaryKey;column:n_id" json:"id"`
	EmployeeID   int    `gorm:"column:n_employee_id" json:"employee_id"`
	Name         string `gorm:"column:c_name" json:"name"`
	NIK          string `gorm:"column:c_nik" json:"nik"`
	Level        string `gorm:"column:c_level" json:"level"`
	LevelID      int    `gorm:"column:n_level_id" json:"level_id"`
	Grade        string `gorm:"column:c_grade" json:"grade"`
	GradeID      int    `gorm:"column:n_grade_id" json:"grade_id"`
	Department   string `gorm:"column:c_department" json:"department"`
	DepartmentID int    `gorm:"column:n_department_id" json:"department_id"`
	// CreatedBy        string           `gorm:"column:c_created_by"`
	// UpdatedBy        string           `gorm:"column:c_updated_by"`
	// CreatedAt        time.Time        `gorm:"default:now();column:d_created_at"`
	// UpdatedAt        time.Time        `gorm:"default:now();column:d_updated_at"`
}

type UserActionData struct {
	ID               int       `gorm:"primaryKey;column:n_id" json:"id"`
	UserID           int       `gorm:"column:n_user_id" json:"user_id"`
	GlobalID         string    `gorm:"column:c_global_id" json:"global_id"`
	IsStartCourse    bool      `gorm:"column:b_isstartcourse" json:"is_startcourse"`
	ModuleAccessed   int       `gorm:"column:n_module_accessed" json:"module_accessed"`
	PretestAccessed  int       `gorm:"column:n_pretest_accessed" json:"pretest_accessed"`
	MateriAccessed   int       `gorm:"column:n_materi_accessed" json:"materi_accessed"`
	PosttestAccessed int       `gorm:"column:n_posttest_accessed" json:"posttest_accessed"`
	CreatedBy        string    `gorm:"column:c_created_by"`
	UpdatedBy        string    `gorm:"column:c_updated_by"`
	CreatedAt        time.Time `gorm:"default:now();column:d_created_at"`
	UpdatedAt        time.Time `gorm:"default:now();column:d_updated_at"`
}
