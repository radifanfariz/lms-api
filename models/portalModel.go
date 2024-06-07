package models

type CompanyData struct {
	BUID             int    `json:"buId"`
	Title            string `json:"title"`
	Name             string `json:"name"`
	AnnualLeaveSetup int    `json:"annualLeaveSetup"`
	CreatedDate      string `json:"createdDate"`
	CreatedBy        string `json:"createdBy"`
	UpdatedDate      string `json:"updatedDate"`
	UpdatedBy        string `json:"updatedBy"`
}

type CompanyDataFromPortal struct {
	Status string        `json:"status"`
	Data   []CompanyData `json:"data"`
}
