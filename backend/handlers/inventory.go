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

// GetUserInventory returns all skins owned by the user
func GetUserInventory(c *gin.Context) {

	// Get user ID from JWT
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	showSold := c.DefaultQuery("show_sold", "false")

	query := database.DB.Preload("Skin").Where("user_id = ?", userID)

	// Conditionally filter out sold items
	if showSold == "false" {
		query = query.Where("is_sold = ?", false)
	}

	// Execute the query and get result 
	var inventory []models.Inventory
	if err := query.Order("acquired_at DESC").Find(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch inventory",
		})
		return
	}

	// calculate statistics
	var totalValue float64
	var itemCount int
	stats := make(map[string]int)

	for _, item := range inventory {
		if !item.IsSold {
			totalValue += item.Value
			itemCount++
			stats[item.Skin.Rarity]++
		}
	}

	// convert to JSON
	var response []map[string]interface{}
	for _, item := range inventory {
		itemData := item.ToJSON()
		itemData["skin"] = item.Skin.ToJSON()
		response = append(response, itemData)
	}

	// send success response 
	c.JSON(http.StatusOK, gin.H{
		"items":        response,
		"total_value": totalValue,
		"item_count":  itemCount,
		"stats":       stats,
	})


	
}

// SellInventoryItem sells a skin for Case Bucks
func SellInventoryItem(c *gin.Context) {
	// Get user ID from JWT
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	// Get item ID from the URl parameter 
	itemID := c.Param("id")
	parsedItemID, err := uuid.Parse(itemID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Invalid item ID",
		})
		return
	}

	// Start database transaction
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}

	}()

	// Fetch item (ensure it belongs to user and isn't already sold)
	var item models.Inventory
	if err := tx.Preload("Skin").Where("id = ? AND user_id = ? AND is_sold = ?", parsedItemID, userID, false ).First(&item).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Item not found or already sold",
		})
		return
	}

	// Get user to update balance
	var user models.User
	if err := tx.First(&user, "id = ?", userID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch user",
		})
		return
	}

	// mark the item as sold 
	item.IsSold = true
	now := time.Now()
	item.SoldAt = &now
	if err := tx.Save(&item).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to mark item as sold",
		})
		return
	}

	// Add Case Bucks to user balance 
	balanceBefore := user.Casebucks
	user.Casebucks += item.Value
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user balance",
		})
		return
	}

	// Create transaction record 
	transaction := models.Transaction{
		UserID:        user.ID,
		Type:          models.TransactionTypeSkinSale,
		Amount:        item.Value,
		BalanceBefore: balanceBefore,
		BalanceAfter:  user.Casebucks,
		Description:   "Sold " + item.Skin.Name,
		ReferenceID:   &parsedItemID,

	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create transaction",
		})
		return
	}

	// Commit transaction (make all changes permanent)
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to complete sale",
		})
		return
	}

	// Send success response 
	c.JSON(http.StatusOK, gin.H{
		"message":       "skin sold successfully!",
		"item":          item.ToJSON(),
		"amount_earned": item.Value,
		"new_balance":   user.Casebucks,
		"transaction_id":   transaction.ID,
	})







}