package models

type EmployeeData struct {
	EmployeeID       string `json:"employeeId"`
	EmployeeNik      string `json:"employeeNik"`
	EmployeeName     string `json:"employeeName"`
	EmployeeStatus   string `json:"employeeStatus"`
	PositionID       string `json:"positionId"`
	EmployeePosition string `json:"employeePosition"`
	DepartmentID     string `json:"departmentId"`
	DepartmentName   string `json:"departmentName"`
	PangkatID        string `json:"pangkatId"`
	PangkatName      string `json:"pangkatName"`
	GradeID          string `json:"gradeId"`
	GradeName        string `json:"gradeName"`
	MainCOmpanyID    string `json:"mainCompanyId"`
	MainCompanyName  string `json:"mainCompanyName"`
	DirectSPV        string `json:"direct_spv"`
	Subordinate      string `json:"subordinate"`
}

type EmployeeDataFromHris struct {
	Message string         `json:"message"`
	Status  bool           `json:"status"`
	Data    []EmployeeData `json:"data"`
}

type GradeData struct {
	GradeID       string `json:"gradeId"`
	MainCompanyID string `json:"mainCompanyId"`
	GradeName     string `json:"gradeName"`
}

type GradeDataFromHris struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    []GradeData `json:"data"`
}

type PositionData struct {
	JabatanID       string `json:"jabatanId"`
	JabatanName     string `json:"jabatanName"`
	MainCompanyID   string `json:"mainCompanyId"`
	MainCompanyName string `json:"mainCompanyName"`
	DepartmentID    string `json:"depatmentId"`
	DepartmentName  string `json:"departmentName"`
	PangkatID       string `json:"pangkatId"`
	PangkatName     string `json:"pangkatName"`
}

type PositionDataFromHris struct {
	Message string         `json:"message"`
	Status  bool           `json:"status"`
	Data    []PositionData `json:"data"`
}
