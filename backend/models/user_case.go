package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserCase tracks purchased cases a user can open later.
type UserCase struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	CaseID    uuid.UUID  `gorm:"type:uuid;not null;index" json:"case_id"`
	IsOpened  bool       `gorm:"not null;default:false" json:"is_opened"`
	OpenedAt  *time.Time `json:"opened_at,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`

	// Relationships
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
	Case Case `gorm:"foreignKey:CaseID;constraint:OnDelete:CASCADE" json:"-"`
}

func (uc *UserCase) BeforeCreate(tx *gorm.DB) error {
	if uc.ID == uuid.Nil {
		uc.ID = uuid.New()
	}
	return nil
}

func (uc *UserCase) ToJSON() map[string]interface{} {
	resp := map[string]interface{}{
		"id":         uc.ID,
		"user_id":    uc.UserID,
		"case_id":    uc.CaseID,
		"is_opened":  uc.IsOpened,
		"created_at": uc.CreatedAt,
		"updated_at": uc.UpdatedAt,
	}
	if uc.OpenedAt != nil {
		resp["opened_at"] = uc.OpenedAt
	}
	return resp
}

