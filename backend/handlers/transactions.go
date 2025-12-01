package handlers

import (
	"net/http"
	"strconv"
	
	
	"github.com/TyronOdame/CS-OPN/backend/database"
	"github.com/TyronOdame/CS-OPN/backend/middleware"
	"github.com/TyronOdame/CS-OPN/backend/models"
	"github.com/gin-gonic/gin"
	
)

// GetUserTransactions returns all transactions for the authenticated user
func GetUserTransactions(c *gin.Context) {
	// Get user ID from JWT
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	// Query parameters for filtering and pagination
	transactionType := c.Query("type") // this filters by type: case_open, skin_sell, skin_buy
	limitStr := c.DefaultQuery("limit", "50")

	// convert limit to integer 
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 50
	}

	// Build the query
	query := database.DB.Where("user_id = ?", userID)

	// filter by transaction type if provided
	if transactionType != "" {
		query = query.Where("type = ?", transactionType)
	}

	// Get transactions from the database
	var transactions []models.Transaction
	if err := query.Order("created_at DESC").Limit(limit).Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch transactions",
		})
		return
	}

	// Convert to JSON
	var response []map[string]interface{}
	for _, tx := range transactions {
		response = append(response, tx.ToJSON())
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": response,
		"count":        len(response),
	})
}