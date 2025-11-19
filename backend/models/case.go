package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Case represents a case item in the database
type Case struct {
	ID          uuid.UUID 	`gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    	`gorm:"not null" json:"name"`
	Price  	    float64   	`gorm:"not null" json:"price"`
	ImageURL    string      `gorm:"not null" json:"image_url"`
	Description string    	`gorm:"type:text" json:"description"`
	IsActive	bool      	`gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

// BeforeCreate hook runs before creating a new skin
func (c *Case) BeforeCreate(tx *gorm.DB) error {
	// Generate a new UUID for the case
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}

// ToJSON converts the Case model to a JSON-compatible map
func (c *Case) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":          c.ID,
		"name":        c.Name,
		"price":       c.Price,
		"image_url":   c.ImageURL,
		"description": c.Description,
		"is_active":   c.IsActive,
		"created_at":  c.CreatedAt,
		"updated_at":  c.UpdatedAt,
	}
}

// CanBeOpened checks if the case is active and can be opened
func (c *Case) CanBeOpened() bool {
	return c.IsActive && c.Price > 0
}