package models

type MateriMetadata struct {
	ID          int    `gorm:"primaryKey;column:n_id" json:"id"`
	ModuleID    int    `gorm:"column:n_module_id" json:"module_id"`
	GlobalID    string `gorm:"column:c_global_id" json:"global_id"`
	Name        string `gorm:"column:c_name" json:"name"`
	Description string `gorm:"column:c_description" json:"description"`
	Src         string `gorm:"column:c_src" json:"src"`
	// CreatedBy   string    `gorm:"column:c_created_by"`
	// UpdatedBy   string    `gorm:"column:c_updated_by"`
	// CreatedAt   time.Time `gorm:"default:now();column:d_created_at"`
	// UpdatedAt   time.Time `gorm:"default:now();column:d_updated_at"`
}

type MateriData struct {
	ID           int    `gorm:"primaryKey;column:n_id" json:"id"`
	ModuleID     int    `gorm:"column:n_module_id" json:"module_id"`
	MateriMetaID int    `gorm:"column:n_materi_meta_id" json:"materi_meta_id"`
	GlobalID     string `gorm:"column:c_global_id" json:"global_id"`
	Type         string `gorm:"column:c_type" json:"type"`
	Src          string `gorm:"column:c_src" json:"src"`
	IsPublished  bool   `gorm:"column:b_ispublished" json:"is_published"`
	// CreatedBy    string         `gorm:"column:c_created_by"`
	// UpdatedBy    string         `gorm:"column:c_updated_by"`
	// CreatedAt    time.Time      `gorm:"default:now();column:d_created_at"`
	// UpdatedAt    time.Time      `gorm:"default:now();column:d_updated_at"`
	Metadata MateriMetadata `gorm:"foreignKey:MateriMetaID" json:"metadata"`
}
