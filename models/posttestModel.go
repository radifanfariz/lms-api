package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type PostTestMetadata struct {
	ID          int       `gorm:"primaryKey;column:n_id" json:"id"`
	ModuleID    int       `gorm:"column:n_module_id" json:"module_id"`
	GlobalID    string    `gorm:"column:c_global_id" json:"global_id"`
	Name        string    `gorm:"column:c_name" json:"name"`
	Description string    `gorm:"column:c_description" json:"description"`
	MaxAccess   int       `gorm:"column:n_max_access" json:"max_access"`
	MinScore    float64   `gorm:"column:n_min_score" json:"min_score"`
	CreatedBy   string    `gorm:"column:c_created_by" json:"created_by"`
	UpdatedBy   string    `gorm:"column:c_updated_by" json:"updated_by"`
	CreatedAt   time.Time `gorm:"default:now();column:d_created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:now();column:d_updated_at" json:"updated_at"`
}

type PostTestData struct {
	ID             int              `gorm:"primaryKey;column:n_id" json:"id"`
	ModuleID       int              `gorm:"column:n_module_id" json:"module_id"`
	PostTestMetaID int              `gorm:"column:n_posttest_meta_id" json:"posttest_meta_id"`
	GlobalID       string           `gorm:"column:c_global_id" json:"global_id"`
	Question       JSONB            `gorm:"column:j_question" json:"question"`
	IsPublished    bool             `gorm:"column:b_ispublished" json:"is_published"`
	CreatedBy      string           `gorm:"column:c_created_by" json:"created_by"`
	UpdatedBy      string           `gorm:"column:c_updated_by" json:"updated_by"`
	CreatedAt      time.Time        `gorm:"column:d_created_at;default:now()" json:"created_at"`
	UpdatedAt      time.Time        `gorm:"column:d_updated_at;default:now()" json:"updated_at"`
	Metadata       PostTestMetadata `gorm:"foreignKey:PostTestMetaID" json:"metadata"`
}

type PostTestResultData struct {
	ID               int              `gorm:"primaryKey;column:n_id" json:"id"`
	UserID           int              `gorm:"column:n_user_id" json:"user_id"`
	GlobalID         string           `gorm:"column:c_global_id" json:"global_id"`
	Score            float64          `gorm:"column:n_score" json:"score"`
	Start            pgtype.Timestamp `gorm:"column:d_start" json:"start"`
	End              pgtype.Timestamp `gorm:"column:d_end" json:"end"`
	Duration         Duration         `gorm:"column:d_duration" json:"duration"`
	Answer           JSONB            `gorm:"type:jsonb;column:j_answer" json:"answer"`
	QuestionAnswered JSONB            `gorm:"type:jsonb;column:j_question_answered" json:"j_question_answered"`
	CreatedBy        string           `gorm:"column:c_created_by" json:"created_by"`
	UpdatedBy        string           `gorm:"column:c_updated_by" json:"updated_by"`
	CreatedAt        time.Time        `gorm:"default:now();column:d_created_at" json:"created_at"`
	UpdatedAt        time.Time        `gorm:"default:now();column:d_updated_at" json:"updated_at"`
}
