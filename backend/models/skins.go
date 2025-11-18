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

// ToJSON converts the Skin model to a JSON-compatible map
func (s *Skin) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":          s.ID,
		"name":        s.Name,
		"weapon_type": s.WeaponType,
		"rarity":      s.Rarity,
		"float":       s.Float,
		"image_url":   s.ImageURL,
		"min_value":   s.MinValue,
		"max_value":   s.MaxValue,
		"description": s.Description,
		"created_at":  s.CreatedAt,
		"updated_at":  s.UpdatedAt,
	}
}

// GetAverageValue calculates the average value of the skin based on its min and max values
func (s *Skin) GetAverageValue() float64 {
	return (s.MinValue + s.MaxValue) / 2
}

//GetRarityColor returns a color code based on the skin's rarity
func (s *Skin) GetRarityColor() string {
	rarityColors := map[string]string{
		"Consumer Grade": "#B0C3D9",
		"Industrial Grade": "#5E98D9",
		"Mil-Spec": "#4B69FF",
		"Restricted": "#8847FF",
		"Classified": "#D32CE6",
		"Covert": "#EB4B4B",
		"Exceedingly Rare": "#FFD700",
	}

	if color, exists := rarityColors[s.Rarity]; exists {
		return color
	}
	return "#FFFFFF" // default to white if rarity not found
}