package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//CaseContent represents the association between a Case and a Skin, including drop chance
type CaseContent struct {
	ID            uuid.UUID     `gorm:"type:uuid;primaryKey" json:"id"`
	CaseID        uuid.UUID     `gorm:"type:uuid;not null;index" json:"case_id"`
	SkinID        uuid.UUID     `gorm:"type:uuid;not null;index" json:"skin_id"`
	DropChance    float64       `gorm:"not null" json:"drop_chance"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`


	// Relationships (GORM will auto-populate these based on foreign keys)
	Case    Case   `gorm:"foreignKey:CaseID;constraint:OnDelete:CASCADE" json:"-"`
	Skin    Skin   `gorm:"foreignKey:SkinID;constraint:OnDelete:CASCADE" json:"-"`
}

//BeforeCreate hook runs before creating a new case content
func (cc *CaseContent) BeforeCreate(tx *gorm.DB) error {
	// Generate UUID if not set
	if cc.ID == uuid.Nil {
		cc.ID = uuid.New()

	}
	return nil
}

//ToJSON converts CaseContent to a JSON-compatible map
func (cc *CaseContent) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":          cc.ID,
		"case_id":    cc.CaseID,
		"skin_id":    cc.SkinID,
		"drop_chance": cc.DropChance,
		"created_at":  cc.CreatedAt,
		"updated_at":  cc.UpdatedAt,
	}
}

//ToJSONWithSkin converts CaseContent to JSON including full skin details
func (cc *CaseContent) ToJSONWithSkin() map[string]interface{}{
	return map[string]interface{}{
		"id":          cc.ID,
		"drop_chance": cc.DropChance,
		"skin":        cc.Skin.ToJSON(),
	}
}

//ToJSONWithCase converts CaseContent to JSON including full case details
func (cc *CaseContent) ToJSONWithCase() map[string]interface{}{
	return map[string]interface{}{
		"id":          cc.ID,
		"drop_chance": cc.DropChance,
		"case":        cc.Case.ToJSON(),
	}
}

//GetDropPercentage returns the drop chance as a percentage
func (cc *CaseContent) GetDropPercentage() float64 {
	return cc.DropChance * 100
}

// IsRare checks this is a rare drop (less than 1% chance)
func (cc *CaseContent) IsRare() bool {
	return cc.DropChance < 0.01
}