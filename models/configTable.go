package models

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by PreTestMetadata to `t_pretest_metadata`
func (PreTestMetadata) TableName() string {
	return "t_pretest_metadata"
}

// TableName overrides the table name used by PreTestData to `t_pretest_data`
func (PreTestData) TableName() string {
	return "t_pretest_data"
}

// TableName overrides the table name used by PreTestResultData to `t_pretest_result_data`
func (PreTestResultData) TableName() string {
	return "t_pretest_result_data"
}

// TableName overrides the table name used by MateriMetadata to `t_materi_metadata`
func (MateriMetadata) TableName() string {
	return "t_materi_metadata"
}

// TableName overrides the table name used by MateriData to `t_materi_data`
func (MateriData) TableName() string {
	return "t_materi_data"
}

// TableName overrides the table name used by MateriData to `t_materi_data`
func (MateriResultData) TableName() string {
	return "t_materi_result_data"
}

// TableName overrides the table name used by PostTestMetadata to `t_posttest_metadata`
func (PostTestMetadata) TableName() string {
	return "t_posttest_metadata"
}

// TableName overrides the table name used by PostTestData to `t_posttest_data`
func (PostTestData) TableName() string {
	return "t_posttest_data"
}

// TableName overrides the table name used by PostTestResultData to `t_posttest_result_data`
func (PostTestResultData) TableName() string {
	return "t_posttest_result_data"
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

// TableName overrides the table name used by User Data to `t_user_action`
func (UserActionData) TableName() string {
	return "t_user_action_data"
}

// TableName overrides the table name used by Access Data to `t_access_data`
func (AccessData) TableName() string {
	return "t_access_data"
}

// TableName overrides the table name used by Gallery Data to `t_gallery_data`
func (GalleryData) TableName() string {
	return "t_gallery_data"
}

// TableName overrides the table name used by Category Data to `t_category_data`
func (CategoryData) TableName() string {
	return "t_category_data"
}

// TableName overrides the table name used by Category Data to `t_certificate_master_data`
func (CertificateMasterData) TableName() string {
	return "t_certificate_master_data"
}

// TableName overrides the table name used by Category Data to `t_certificate_user_data`
func (CertificateUserData) TableName() string {
	return "t_certificate_user_data"
}
