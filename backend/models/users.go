package models

import(
	"time"


	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//This will be the user model for the application
type User struct {
	ID			uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	Email   	string         `gorm:"uniqueIndex;not null" json:"email"`
	Username 	string        `gorm:"uniqueIndex;not null" json:"username"`
	Password    string        `gorm:"not null" json:"-"`
	Casebucks   float64       `gorm:"default:0" json:"casebucks"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Inventory   []InventoryItem `gorm:"foreignKey:UserID" json:"-"`
	Transactions []Transaction  `gorm:"foreignKey:UserID" json:"-"`
}

// function to handle pre-user creation
func (u *User) 