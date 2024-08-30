package models

import "time"

type ModuleMetadata struct {
	ID              int       `gorm:"primaryKey;column:n_id" json:"id"`
	GlobalID        string    `gorm:"column:c_global_id" json:"global_id"`
	Name            string    `gorm:"column:c_name" json:"name"`
	Description     string    `gorm:"column:c_description" json:"description"`
	Src             string    `gorm:"column:c_src" json:"src"`
	LearningJourney string    `gorm:"column:c_learning_journey" json:"learning_journey"`
	Category        string    `gorm:"column:c_category" json:"category"`
	MaxMonth        int       `gorm:"column:n_max_month" json:"max_month"`
	Seq             int       `gorm:"column:n_seq" json:"seq"`
	CreatedBy       string    `gorm:"column:c_created_by" json:"created_by"`
	UpdatedBy       string    `gorm:"column:c_updated_by" json:"updated_by"`
	CreatedAt       time.Time `gorm:"column:d_created_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:d_updated_at" json:"updated_at"`
}

/* this relation with GlobalID --> */
type ModuleData struct {
	ID               int              `gorm:"primaryKey;column:n_id" json:"id"`
	GlobalID         string           `gorm:"column:c_global_id" json:"global_id"`
	ModuleMetaID     int              `gorm:"column:n_module_meta_id" json:"module_meta_id"`
	PreTestMetaID    int              `gorm:"column:n_pretest_meta_id" json:"pretest_meta_id"`
	PreTestID        int              `gorm:"column:n_pretest_id" json:"pretest_id"`
	MateriMetaID     int              `gorm:"column:n_materi_meta_id" json:"materi_meta_id"`
	MateriID         int              `gorm:"column:n_materi_id" json:"materi_id"`
	PostTestMetaID   int              `gorm:"column:n_posttest_meta_id" json:"posttest_meta_id"`
	PostTestID       int              `gorm:"column:n_posttest_id" json:"posttest_id"`
	UserID           int              `gorm:"column:n_user_id" json:"user_id"`
	GradeID          int              `gorm:"column:n_grade_id" json:"grade_id"`
	Metadata         ModuleMetadata   `gorm:"foreignKey:GlobalID;references:GlobalID" json:"metadata"`
	UserData         UserData         `gorm:"foreignKey:UserID" json:"user_data"`
	PreTestMetadata  PreTestMetadata  `gorm:"foreignKey:GlobalID;references:GlobalID" json:"pretest_metadata"`
	MateriMetadata   MateriMetadata   `gorm:"foreignKey:GlobalID;references:GlobalID" json:"materi_metadata"`
	PostTestMetadata PostTestMetadata `gorm:"foreignKey:GlobalID;references:GlobalID" json:"posttest_metadata"`
	PreTestData      PreTestData      `gorm:"foreignKey:GlobalID;references:GlobalID" json:"pretest_data"`
	MateriData       []MateriData     `gorm:"foreignKey:GlobalID;references:GlobalID" json:"materi_data"`
	PostTestData     PostTestData     `gorm:"foreignKey:GlobalID;references:GlobalID" json:"posttest_data"`
	AccessData       []AccessData     `gorm:"foreignKey:GlobalID;references:GlobalID" json:"access_data"`
	IsPublished      *bool            `gorm:"column:b_ispublished" json:"is_published"`
	CreatedBy        string           `gorm:"column:c_created_by" json:"created_by"`
	UpdatedBy        string           `gorm:"column:c_updated_by" json:"updated_by"`
	CreatedAt        time.Time        `gorm:"column:d_created_at" json:"created_at"`
	UpdatedAt        time.Time        `gorm:"column:d_updated_at" json:"updated_at"`
}

/* this relation with Each ID (ModuleMetaID, PreTestMetaID, etc) --> */
// type ModuleData struct {
// 	ID               int              `gorm:"primaryKey;column:n_id" json:"id"`
// 	GlobalID         string           `gorm:"column:c_global_id" json:"global_id"`
// 	ModuleMetaID     int              `gorm:"column:n_module_meta_id" json:"module_meta_id"`
// 	PreTestMetaID    int              `gorm:"column:n_pretest_meta_id" json:"pretest_meta_id"`
// 	PreTestID        int              `gorm:"column:n_pretest_id" json:"pretest_id"`
// 	MateriMetaID     int              `gorm:"column:n_materi_meta_id" json:"materi_meta_id"`
// 	MateriID         int              `gorm:"column:n_materi_id" json:"materi_id"`
// 	PostTestMetaID   int              `gorm:"column:n_posttest_meta_id" json:"posttest_meta_id"`
// 	PostTestID       int              `gorm:"column:n_posttest_id" json:"posttest_id"`
// 	UserID           int              `gorm:"column:n_user_id" json:"user_id"`
// 	GradeID          int              `gorm:"column:n_grade_id" json:"grade_id"`
// 	Metadata         ModuleMetadata   `gorm:"foreignKey:ModuleMetaID" json:"metadata"`
// 	UserData         UserData         `gorm:"foreignKey:UserID" json:"user_data"`
// 	PreTestMetadata  PreTestMetadata  `gorm:"foreignKey:PreTestMetaID" json:"pretest_metadata"`
// 	MateriMetadata   MateriMetadata   `gorm:"foreignKey:MateriMetaID" json:"materi_metadata"`
// 	PostTestMetadata PostTestMetadata `gorm:"foreignKey:PostTestMetaID" json:"posttest_metadata"`
// 	PreTestData      PreTestData      `gorm:"foreignKey:PreTestID" json:"pretest_data"`
// 	MateriData       MateriData       `gorm:"foreignKey:MateriID" json:"materi_data"`
// 	PostTestData     PostTestData     `gorm:"foreignKey:PostTestID" json:"posttest_data"`
// }
