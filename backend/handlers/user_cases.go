package handlers

import (
	"net/http"
	"time"

	"github.com/TyronOdame/CS-OPN/backend/database"
	"github.com/TyronOdame/CS-OPN/backend/middleware"
	"github.com/TyronOdame/CS-OPN/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetUserCases returns unopened cases the user has purchased.
func GetUserCases(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var userCases []models.UserCase
	if err := database.DB.
		Preload("Case").
		Where("user_id = ? AND is_opened = ?", userID, false).
		Order("created_at DESC").
		Find(&userCases).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch purchased cases"})
		return
	}

	resp := make([]map[string]interface{}, 0, len(userCases))
	for _, uc := range userCases {
		item := uc.ToJSON()
		item["case"] = uc.Case.ToJSON()
		resp = append(resp, item)
	}

	c.JSON(http.StatusOK, gin.H{
		"cases": resp,
		"count": len(resp),
	})
}

// OpenPurchasedCase opens a bought case from user's case inventory.
func OpenPurchasedCase(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userCaseID := c.Param("id")
	parsedUserCaseID, err := uuid.Parse(userCaseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchased case ID"})
		return
	}

	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var userCase models.UserCase
	if err := tx.Preload("Case").Where("id = ? AND user_id = ? AND is_opened = ?", parsedUserCaseID, userID, false).First(&userCase).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchased case not found or already opened"})
		return
	}

	var contents []models.CaseContent
	if err := tx.Preload("Skin").Where("case_id = ?", userCase.CaseID).Find(&contents).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch case contents"})
		return
	}

	selectedContent := selectRandomSkin(contents)
	if selectedContent == nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select skin"})
		return
	}

	randomFloat := generateFloat()
	skin := selectedContent.Skin
	floatMultiplier := 1.0 - randomFloat
	skinValue := skin.MinValue + (skin.MaxValue-skin.MinValue)*floatMultiplier

	inventory := models.Inventory{
		UserID:       userID,
		SkinID:       skin.ID,
		Float:        randomFloat,
		AcquiredFrom: userCase.Case.Name,
		Value:        skinValue,
		IsSold:       false,
	}
	if err := tx.Create(&inventory).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add skin to inventory"})
		return
	}

	now := time.Now()
	userCase.IsOpened = true
	userCase.OpenedAt = &now
	if err := tx.Save(&userCase).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark case as opened"})
		return
	}

	transaction := models.Transaction{
		UserID:        userID,
		Type:          models.TransactionTypeCaseOpen,
		Amount:        0,
		BalanceBefore: 0,
		BalanceAfter:  0,
		Description:   "Opened purchased " + userCase.Case.Name,
		ReferenceID:   &userCase.CaseID,
	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	var user models.User
	if err := tx.First(&user, "id = ?", userID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}
	transaction.BalanceBefore = user.Casebucks
	transaction.BalanceAfter = user.Casebucks
	if err := tx.Save(&transaction).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update transaction"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to complete case opening"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "Case opened successfully!",
		"case":           userCase.Case.ToJSON(),
		"skin":           skin.ToJSON(),
		"float":          randomFloat,
		"condition":      inventory.GetCondition(),
		"value":          skinValue,
		"new_balance":    user.Casebucks,
		"inventory_id":   inventory.ID,
		"transaction_id": transaction.ID,
	})
}

