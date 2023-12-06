package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type PreTestMetadata struct {
	ID          int       `gorm:"primaryKey;column:n_id"`
	ModuleID    int       `gorm:"column:n_module_id"`
	Slug        string    `gorm:"column:c_slug"`
	Name        string    `gorm:"column:c_name"`
	Description string    `gorm:"column:c_description"`
	Src         string    `gorm:"column:c_src"`
	MaxAccess   int       `gorm:"column:n_max_access"`
	MinScore    float64   `gorm:"column:n_min_score"`
	CreatedBy   string    `gorm:"column:c_created_by"`
	UpdatedBy   string    `gorm:"column:c_updated_by"`
	CreatedAt   time.Time `gorm:"default:now();column:d_created_at"`
	UpdatedAt   time.Time `gorm:"default:now();column:d_updated_at"`
}

type PreTestData struct {
	ID            int       `gorm:"primaryKey;column:n_id"`
	ModuleID      int       `gorm:"column:n_module_id"`
	PreTestMetaID int       `gorm:"column:n_pretest_meta_id"`
	Slug          string    `gorm:"column:c_slug"`
	Question      JSONB     `gorm:"type:jsonb;column:j_question"`
	IsPublished   bool      `gorm:"column:b_ispublished"`
	CreatedBy     string    `gorm:"column:c_created_by"`
	UpdatedBy     string    `gorm:"column:c_updated_by"`
	CreatedAt     time.Time `gorm:"default:now();column:d_created_at"`
	UpdatedAt     time.Time `gorm:"default:now();column:d_updated_at"`
}

type PreTestResultData struct {
	ID               int              `gorm:"primaryKey;column:n_id"`
	UserID           int              `gorm:"column:n_user_id"`
	Score            float64          `gorm:"column:n_score"`
	Start            pgtype.Timestamp `gorm:"column:d_start"`
	End              pgtype.Timestamp `gorm:"column:d_end"`
	Duration         Duration         `gorm:"column:d_duration"`
	Answer           JSONB            `gorm:"type:jsonb;column:j_answer"`
	QuestionAnswered JSONB            `gorm:"type:jsonb;column:j_question_answered"`
	CreatedBy        string           `gorm:"column:c_created_by"`
	UpdatedBy        string           `gorm:"column:c_updated_by"`
	CreatedAt        time.Time        `gorm:"default:now();column:d_created_at"`
	UpdatedAt        time.Time        `gorm:"default:now();column:d_updated_at"`
}

// // Assuming you have a GORM DB instance named "db"
// db.AutoMigrate(&PretestMetadata{})
