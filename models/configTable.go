package models

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by PretestMeatadata to `t_pretest_metadata`
func (PreTestMetadata) TableName() string {
	return "t_pretest_metadata"
}

// TableName overrides the table name used by Pretestdata to `t_pretest_data`
func (PreTestData) TableName() string {
	return "t_pretest_data"
}

// TableName overrides the table name used by PretestResultData to `t_pretest_result_data`
func (PreTestResultData) TableName() string {
	return "t_pretest_result_data"
}

// TableName overrides the table name used by ModuleMetadata to `t_module_metadata`
func (ModuleMetadata) TableName() string {
	return "t_module_metadata"
}

// TableName overrides the table name used by ModuleMetadata to `t_module_metadata`
func (ModuleData) TableName() string {
	return "t_module_data"
}

// TableName overrides the table name used by User Data to `t_user_data`
func (UserData) TableName() string {
	return "t_user_data"
}
