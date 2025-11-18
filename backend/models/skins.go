package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

)

// Skin represents a skin item in the database
type Skin struct {
	ID    		uuid.UUID    `gorm:"type:uuid;primaryKey" json:"id"`
	Name  		string       `gorm:"not null" json:"name"`
	WeaponType  string	     `gorm:"not null" json:"weapon_type"`	
	Rarity      string       `gorm:"not null" json:"rarity"`
	Float       float64      `gorm:"default:0.5" json:"float"`
	ImageURL    string       `gorm:"not null" json:"image_url"`
	MinValue    float64      `gorm:"not null" json:"min_value"`
	MaxValue    float64      `gorm:"not null" json:"max_value"`
	Description string       `gorm:"type:text" json:"description"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

// BeforeCreate is a GORM hook that is triggered before a new Skin record is created
func (s *Skin) BeforeCreate(tx *gorm.DB) error {
	// Generate a new UUID for the skin
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}