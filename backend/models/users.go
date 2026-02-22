package models

import (
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
	LastDailyRewardAt *time.Time `json:"last_daily_reward_at,omitempty"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
// 	Inventory   []InventoryItem `gorm:"foreignKey:UserID" json:"-"`
// 	Transactions []Transaction  `gorm:"foreignKey:UserID" json:"-"`
}

// function to handle pre-user creation
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	// Users start with 100 CBs
	if u.Casebucks == 0 {
		u.Casebucks = 100
	}
	return nil
}

// Hashing the users password before saving to the database
func (u *User) HashPassword(password string) error {

	// This is where bcrypt generates the hashed password with the users given password(the byte conversion is necessary)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Checking to see if the password given matches the hashed password stored in the db
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// This function returns a safe JSON version of a user without any of the sensitive information
func (u *User) ToJSON() map[string]interface{} {
	return map[string]interface{} {
		"id":         u.ID,
		"email":      u.Email,
		"username":   u.Username,
		"casebucks":  u.Casebucks,
		"last_daily_reward_at": u.LastDailyRewardAt,
		"created_at": u.CreatedAt,
		"updated_at": u.UpdatedAt,
	}
}
