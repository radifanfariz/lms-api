package models

import "time"

type UserData struct {
	ID                *int   `gorm:"primaryKey;column:n_id" json:"id"`
	EmployeeID        int    `gorm:"column:n_employee_id" json:"employee_id"`
	Name              string `gorm:"column:c_name" json:"name"`
	NIK               string `gorm:"column:c_nik" json:"nik"`
	MainCompany       string `gorm:"column:c_main_company" json:"main_company"`
	MainCompanyID     int    `gorm:"column:n_main_company_id" json:"main_company_id"`
	Level             string `gorm:"column:c_level" json:"level"`
	LevelID           *int   `gorm:"column:n_level_id" json:"level_id"`
	Grade             string `gorm:"column:c_grade" json:"grade"`
	GradeID           int    `gorm:"column:n_grade_id" json:"grade_id"`
	Department        string `gorm:"column:c_department" json:"department"`
	DepartmentID      int    `gorm:"column:n_department_id" json:"department_id"`
	LearningJourney   string `gorm:"column:c_learning_journey" json:"learning_journey"`
	LearningJourneyID int    `gorm:"column:n_learning_journey_id" json:"learning_journey_id"`
	Role              string `gorm:"column:c_role" json:"role"`
	RoleID            int    `gorm:"column:n_role_id" json:"role_id"`
	Status            string `gorm:"column:c_status" json:"status"`
	StatusID          int    `gorm:"column:n_status_id" json:"status_id"`
	IsActive          *bool  `gorm:"column:b_isactive" json:"is_active"`
	Position          string `gorm:"column:c_position" json:"position"`
	PositionID        int    `gorm:"column:n_position_id" json:"position_id"`
	AlternativeID     string `gorm:"column:c_alternative_id" json:"alternative_id"`
	// Password          string    `gorm:"column:c_password" json:"password"`
	CreatedBy string    `gorm:"column:c_created_by" json:"created_by"`
	UpdatedBy string    `gorm:"column:c_updated_by" json:"updated_by"`
	CreatedAt time.Time `gorm:"default:now();column:d_created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:now();column:d_updated_at" json:"updated_at"`
}

type UserActionData struct {
	ID               int        `gorm:"primaryKey;column:n_id" json:"id"`
	UserID           int        `gorm:"column:n_user_id" json:"user_id"`
	GlobalID         string     `gorm:"column:c_global_id" json:"global_id"`
	IsStartCourse    *bool      `gorm:"column:b_isstartcourse" json:"is_startcourse"`
	ModuleAccessed   int        `gorm:"column:n_module_accessed" json:"module_accessed"`
	PretestAccessed  int        `gorm:"column:n_pretest_accessed" json:"pretest_accessed"`
	MateriAccessed   int        `gorm:"column:n_materi_accessed" json:"materi_accessed"`
	PosttestAccessed int        `gorm:"column:n_posttest_accessed" json:"posttest_accessed"`
	CreatedBy        string     `gorm:"column:c_created_by" json:"created_by"`
	UpdatedBy        string     `gorm:"column:c_updated_by" json:"updated_by"`
	CreatedAt        time.Time  `gorm:"default:now();column:d_created_at" json:"created_at"`
	UpdatedAt        time.Time  `gorm:"default:now();column:d_updated_at" json:"updated_at"`
	UserData         UserData   `gorm:"foreignKey:UserID" json:"user_data"`
	ModuleData       ModuleData `gorm:"foreignKey:GlobalID;references:GlobalID" json:"module_data"`
}
