package models

import (
    "time"

    "github.com/google/uuid"
)

// CaseSkin represents the many-to-many relationship between Cases and Skins
// Each case can contain multiple skins, each with a specific drop rate
type CaseSkin struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    CaseID    uuid.UUID `gorm:"type:uuid;not null" json:"case_id"`
    SkinID    uuid.UUID `gorm:"type:uuid;not null" json:"skin_id"`
    DropRate  float64   `gorm:"not null" json:"drop_rate"` // Percentage chance (0-100)
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`

    // Relationships
    Case *Case `gorm:"foreignKey:CaseID;constraint:OnDelete:CASCADE" json:"case,omitempty"`
    Skin *Skin `gorm:"foreignKey:SkinID;constraint:OnDelete:CASCADE" json:"skin,omitempty"`
}

// TableName specifies the table name for GORM
func (CaseSkin) TableName() string {
    return "case_skins"
}