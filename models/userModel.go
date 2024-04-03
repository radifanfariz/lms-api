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

/*SSO User Data Model*/
type UserDataSSO struct {
	UserToken   string `json:"userToken"`
	NIK         string `json:"nik"`
	LastLogin   string `json:"lastLogin"`
	CreatedDate string `json:"createdDate"`
	UpdatedBy   string `json:"updatedBy"`
	CreatedBy   string `json:"createdBy"`
	Name        string `json:"name"`
	UpdatedDate string `json:"updatedDate"`
	UserId      string `json:"userId"`
	Status      string `json:"status"`
}

/* Portal User Data Model */
type Token struct {
	User string `json:"user"`
	SSO  string `json:"sso"`
}
type UserBU struct {
	BUID             string `json:"buId"`
	Title            string `json:"title"`
	Name             string `json:"name"`
	AnnualLeaveSetup string `json:"annualLeaveSetup"`
	CreatedDate      string `json:"createdDate"`
	CreatedBy        string `json:"createdBy"`
	UpdatedDate      string `json:"updatedDate"`
	UpdatedBy        string `json:"updatedBy"`
}
type Menus struct {
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	MenuID string `json:"menuId"`
	Seq    string `json:"seq"`
}
type UserDataPortal struct {
	SSOID        string   `json:"ssoId"`
	HRISID       int      `json:"hrisId"`
	WorkScheme   int      `json:"workScheme"`
	LastLogin    string   `json:"lastLogin"`
	Access       [][]int  `json:"access"`
	Gender       string   `json:"gender"`
	Roles        string   `json:"roles"`
	UpdatedDate  string   `json:"updatedDate"`
	Office       []string `json:"office"`
	LastNotified int      `json:"lastNotified"`
	UserID       int      `json:"userId"`
	BirthDate    string   `json:"birthDate"`
	Token        []Token  `json:"token"`
	NIK          string   `json:"nik"`
	HRDDoc       int      `json:"hrdDoc"`
	JoinDate     string   `json:"joinDate"`
	BU           [][]any  `json:"bu"`
	UserBU       UserBU   `json:"userBu"`
	Name         string   `json:"name"`
	Menus        []Menus  `json:"menus"`
	Email        string   `json:"email"`
	Status       string   `json:"status"`
}
