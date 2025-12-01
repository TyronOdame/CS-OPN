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

//GetUserInventory returns all skins owned by the user
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







}