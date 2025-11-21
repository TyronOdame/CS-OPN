package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

)

// Inventory represents a skin owned by a user
type Inventory struct {
	ID              uuid.UUID    `gorm:"type:uuid;primaryKey" json:"id"`
	UserID          uuid.UUID    `gorm:"type:uuid;not null;index" json:"user_id"`
	SkinID          uuid.UUID    `gorm:"type:uuid;not null;index" json:"skin_id"`
	Float           float64      `gorm:"not null" json:"float"`
	AcquiredFrom    string       `gorm:"not null" json:"acquired_from"`
	Value           float64      `gorm:"not null" json:"value"`
	IsSold          bool         `gorm:"not null;default:false" json:"is_sold"`
	SoldAt          *time.Time   `json:"sold_at"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
	
	

	// Relationships
	User User `gorm:"foreignKy:UserID;constraint:OnDelete:CASCADE" json:"-"`
	Skin Skin `gorm:"foreignKey:SkinID;constraint:onDelete:CASCADE" json:"-"`
}

// BeforeCreate hook runs before creating a new inventory item
func (i *Inventory) BeforeCreate(tx *gorm.DB) error {
	// Generate UUID if not set
	if i.ID == uuid.Nil {
		i.ID = uuid.New()

	}
	return nil
}


// ToJSON concerts Inventory to a JSON-safe map
func (i *Inventory) ToJSON() map[string] interface{} {
	response := map[string]interface{}{
		"id":           i.ID,
		"user_id":      i.UserID,
		"skin_id":      i.SkinID,
		"float":        i.Float,
		"acquired_from": i.AcquiredFrom,
		"value":        i.Value,
		"is_sold":      i.IsSold,
		"created_at":   i.CreatedAt,
		"updated_at":   i.UpdatedAt,
	}

	// only include sold_at if the item has been sold
	if i.SoldAt != nil {
		response["sold_at"] = i.SoldAt
	}
	
	return response
}

// ToJSONWithSkin converts Inventory to JSON including full skin details
func (i *Inventory) ToJSONWithSkin() map[string] interface{} {
	response := i.ToJSON()
	response["skin"] = i.Skin.ToJSON()
	return response
}

// GetCondition returns the war condition base on the float value
func (i *Inventory) GetCondition() string {
	if i.Float < 0.07 {
		return "Factory New"
	} else if i.Float < 0.15 {
		return "Minimal Wear"
	} else if i.Float < 0.38 {
		return "Field-Tested"
	} else if i.Float < 0.45 {
		return "Well-Worn"
	} else {
		return "Battle-Scarred"
	}
}

// CanBeSold checks if the inventory item can be sold
func(i *Inventory) CanBeSold() bool {
	return !i.IsSold
}

//sell marks the inventory item as sold at the given time
func (i *Inventory) Sell(tx *gorm.DB) error {
	now := time.Now()
	i.IsSold = true
	i.SoldAt = &now
	return tx.Save(i).Error
}