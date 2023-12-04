package models

type UserData struct {
	ID           int    `gorm:"primaryKey;column:n_id"`
	EmployeeID   int    `gorm:"column:n_employee_id"`
	Name         string `gorm:"column:c_name"`
	NIK          string `gorm:"column:c_nik"`
	Level        string `gorm:"column:c_level"`
	LevelID      int    `gorm:"column:n_level_id"`
	Grade        string `gorm:"column:c_grade"`
	GradeID      int    `gorm:"column:n_grade_id"`
	Department   string `gorm:"column:c_department"`
	DepartmentID int    `gorm:"column:n_department_id"`
	// CreatedBy        string           `gorm:"column:c_created_by"`
	// UpdatedBy        string           `gorm:"column:c_updated_by"`
	// CreatedAt        time.Time        `gorm:"default:now();column:d_created_at"`
	// UpdatedAt        time.Time        `gorm:"default:now();column:d_updated_at"`
}
