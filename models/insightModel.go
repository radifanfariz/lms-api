package models

type TotalInsightsData struct {
	TotalUsers                    int `gorm:"column:n_total_users" json:"total_users"`
	TotalLearningModules          int `gorm:"column:n_total_learning_modules" json:"total_learning_modules"`
	TotalPretestParticipants      int `gorm:"column:n_total_pretest_participants" json:"total_pretest_participants"`
	TotalMateriParticipants       int `gorm:"column:n_total_materi_participants" json:"total_materi_participants"`
	TotalPosttestParticipants     int `gorm:"column:n_total_posttest_participants" json:"total_posttest_participants"`
	TotalEnrolledLearningModules  int `gorm:"column:n_total_enrolled_learning_modules" json:"total_enrolled_learning_modules"`
	TotalGraduatedLearningModules int `gorm:"column:n_total_graduated_learning_modules" json:"total_graduated_learning_modules"`
}

type TotalEnrolledPerMonthInsightsData struct {
	Date          string `gorm:"column:c_date" json:"date"`
	TotalEnrolled int    `gorm:"column:n_total_enrolled" json:"total_enrolled"`
}

type EnrolledInsightsData struct {
	GlobalID      string `gorm:"column:c_global_id" json:"global_Id"`
	ModuleName    string `gorm:"column:c_module_name" json:"module_name"`
	TotalEnrolled int    `gorm:"column:n_total_enrolled" json:"total_enrolled"`
}

type UserInsightsData struct {
	ID               int     `gorm:"column:n_id" json:"id"`
	NIK              string  `gorm:"column:c_nik" json:"nik"`
	Name             string  `gorm:"column:c_name" json:"name"`
	LearningJourney  string  `gorm:"column:c_learning_journey" json:"learning_journey"`
	ModuleName       string  `gorm:"column:c_module_name" json:"module_name"`
	PretestScore     float64 `gorm:"column:n_pretest_score" json:"pretest_score"`
	PosttestScore    float64 `gorm:"column:n_posttest_score" json:"posttest_score"`
	PosttestMinScore float64 `gorm:"column:n_posttest_min_score" json:"posttest_min_score"`
}
