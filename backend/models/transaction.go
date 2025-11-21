package models


import (
	"time"


	"github.com/google/uuid"
	"gorm.io/gorm"
)

//TransactionType defines the type of transaction
type TransactionType string 

const (
	TransactionTypeSkinSell      TransactionType = "skin_sell"
	TransactionTypeCaseOpen      TransactionType = "case_open"
	TransactionTypeDailyLogin    TransactionType = "daily_login"
	TransactionTypeRegistration  TransactionType = "registration"
	TransactionTypeRefund        TransactionType = "refund"
)

// Transaction represents a CaseBucks transaction
type Transaction struct {
	ID             uuid.UUID        `gorm:"type:uuid;primaryKey" json:"id"`
	UserID         uuid.UUID        `gorm:"type:uuid;not null;index" json:"user_id"`
	Type           TransactionType  `gorm:"type:varchar(50);not null" json:"type"`
	Amount         float64          `gorm:"not null" json:"amount"`
	BalanceBefore  float64          `gorm:"not null" json:"balance_before"`
	BalanceAfter   float64          `gorm:"not null" json:"balance_after"`
	Description    string           `gorm:"type:text" json:"description"`
	ReferenceID    *uuid.UUID       `gorm:"type:uuid;index" json:"reference_id,omitempty"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`


	// Relationships
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`

}

// BeforeCreate hook runs before creating a new transaction
func (t *Transaction) BeforeCreate(tx *gorm.DB) error {
	// Generate UUID if not set
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}

// ToJSON converts Transaction to a JSON-compatible map
func (t *Transaction) ToJSON() map[string]interface{} {
	response := map[string]interface{}{
		"id":             t.ID,
		"user_id":        t.UserID,
		"type":           t.Type,
		"amount":         t.Amount,
		"balance_before": t.BalanceBefore,
		"balance_after":  t.BalanceAfter,
		"description":    t.Description,
		"created_at":     t.CreatedAt,
		"updated_at":     t.UpdatedAt,
	}

	// Only include reference_id if it exists
	if t.ReferenceID != nil {
		response["reference_id"] = t.ReferenceID
	}
	return response
}

//IsDebit checks if this transaction reduced the balance 
func (t *Transaction) IsDebit() bool {
	return t.Amount < 0
}

// IsCredit checks if this transaction increased the balance
func (t *Transaction) IsCredit() bool {
	return t.Amount > 0
}

// GetAbsoluteAmount returns the absolute value of the transaction amount 
func (t *Transaction) GetAbsoluteAmount() float64 {
	if t.Amount < 0 {
		return -t.Amount
	}
	return t.Amount
}