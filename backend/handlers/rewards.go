package handlers

import (
	"time"

	"github.com/TyronOdame/CS-OPN/backend/database"
	"github.com/TyronOdame/CS-OPN/backend/models"
	"gorm.io/gorm"
)

const DailyLoginRewardAmount = 100.0

// applyDailyLoginReward grants 100 CB once every 24h on login.
func applyDailyLoginReward(tx *gorm.DB, user *models.User) (bool, error) {
	now := time.Now()
	if user.LastDailyRewardAt != nil && now.Sub(*user.LastDailyRewardAt) < 24*time.Hour {
		return false, nil
	}

	balanceBefore := user.Casebucks
	user.Casebucks += DailyLoginRewardAmount
	user.LastDailyRewardAt = &now
	if err := tx.Save(user).Error; err != nil {
		return false, err
	}

	transaction := models.Transaction{
		UserID:        user.ID,
		Type:          models.TransactionTypeDailyLogin,
		Amount:        DailyLoginRewardAmount,
		BalanceBefore: balanceBefore,
		BalanceAfter:  user.Casebucks,
		Description:   "Daily login reward",
	}
	if err := tx.Create(&transaction).Error; err != nil {
		return false, err
	}

	return true, nil
}

// ApplyDailyRewardForUser is used by non-auth login flows.
func ApplyDailyRewardForUser(user *models.User) (bool, error) {
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	claimed, err := applyDailyLoginReward(tx, user)
	if err != nil {
		tx.Rollback()
		return false, err
	}
	if err := tx.Commit().Error; err != nil {
		return false, err
	}
	return claimed, nil
}

