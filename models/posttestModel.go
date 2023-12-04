package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type PostTestMetadata struct {
	ID          int       `gorm:"primaryKey;column:n_id"`
	ModuleID    int       `gorm:"column:n_module_id"`
	Slug        string    `gorm:"column:c_slug"`
	Name        string    `gorm:"column:c_name"`
	Description string    `gorm:"column:c_description"`
	MaxAccess   int       `gorm:"column:n_max_access"`
	MinScore    float64   `gorm:"column:n_min_score"`
	CreatedBy   string    `gorm:"column:c_created_by"`
	UpdatedBy   string    `gorm:"column:c_updated_by"`
	CreatedAt   time.Time `gorm:"default:now();column:d_created_at"`
	UpdatedAt   time.Time `gorm:"default:now();column:d_updated_at"`
}

type PostTestData struct {
	ID             int       `gorm:"primaryKey;column:n_id"`
	ModuleID       int       `gorm:"column:n_module_id"`
	PostTestMetaID int       `gorm:"column:n_posttest_meta_id"`
	Slug           string    `gorm:"column:c_slug"`
	Question       JSONB     `gorm:"column:j_question"`
	CreatedBy      string    `gorm:"column:c_created_by"`
	UpdatedBy      string    `gorm:"column:c_updated_by"`
	CreatedAt      time.Time `gorm:"column:d_created_at;default:now()"`
	UpdatedAt      time.Time `gorm:"column:d_updated_at;default:now()"`
}

type PostTestResultData struct {
	ID               int              `gorm:"primaryKey"`
	UserID           int              `gorm:"column:n_user_id"`
	Score            float64          `gorm:"column:n_score"`
	Start            pgtype.Timestamp `gorm:"column:d_start"`
	End              pgtype.Timestamp `gorm:"column:d_end"`
	Duration         pgtype.Interval  `gorm:"column:d_duration"`
	Answer           JSONB            `gorm:"type:jsonb;column:j_answer"`
	QuestionAnswered JSONB            `gorm:"type:jsonb;column:j_question_answered"`
	CreatedBy        string           `gorm:"column:c_created_by"`
	UpdatedBy        string           `gorm:"column:c_updated_by"`
	CreatedAt        time.Time        `gorm:"default:now();column:d_created_at"`
	UpdatedAt        time.Time        `gorm:"default:now();column:d_updated_at"`
}
