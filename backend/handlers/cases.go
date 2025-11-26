package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/TyronOdame/CS-OPN/backend/database"
	"github.com/TyronOdame/CS-OPN/backend/middleware"
	"github.com/TyronOdame/CS-OPN/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//GetAllCases returns all active cases
func GetAllCases(c *gin.Context) {
	var cases []models.Case

	// Only get active cases
	if err := database.DB.Where("is_active = ?", true).Find(&cases).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch cases",
		})
		return
	}

	// Convert to JSON
	var response []map[string]interface{}
	for _, cs := range cases {
		response = append(response, cs.ToJSON())
	}

	c.JSON(http.StatusOK, gin.H{
		"cases": response,
	})
}

//GetCaseByID returns a case with all possible skins
func GetCaseByID(c *gin.Context) {
	caseID = c.Param("id")

	// Validate UUID
	if _, err := uuid.Parse(caseID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid case ID",
		})
		return
	}
		
	var caseItem models.Case
	if err := database.DB.First(&caseItem, "id = ? AND is_active = ?", caseID, true).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Case not found",
		})
		return
	}

	// get all skins in this case with drop chances
	var contents []models.CaseContent
	if err := database.DB.Preload("Skin").Where("case_id = ?", caseID).Find(&contents).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Case not found",
			
		})
		return
	}

	// build response with case info and skins
	var skins []map[string]interface{}
	for_, content := range contents {
		skinData := content.Skin.ToJSON()
		skinData["drop_chance"] = content.DropChance
		skinData["drop_percentage"] = content.GetDropPercentage()
		skins = append(skins, skinData)
	}

	response := caseItem.ToJSON()
	response["skins"] = skins

	c.JSON(http.StatusOk, response)

	

	
}

// OpenCase handles case opening logic
func OpenCase(c *gin.Context) {

	// get case ID from JWT
	userID, exists := middleware.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	// Get case ID from URL
	caseID := c.Param("id")
	parsedCaseID, err := uuid.Parse(caseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid case ID",
		})
		return
	}

	// start transaction
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// get case details
	var caseItem models.Case
	if err := tx.First(&caseItem, "id = ?", parsedCaseID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Case not found",
		})
		return
	}

	//check if case is active
	if !caseItem.IsActive {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Case is no longer available",
		})
		return
	}

	//get user and check balance
	var user models.User
	if err := tx.First(&user, "id = ?", userID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "User not found",
		})
		return
	}

	// check if user has enough balance
	if user.Balance < caseItem.Price {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Insufficient balance",
		})
		return
	}

	//Get all possible skins in the case with drop chances
	var contents []models.CaseContent
	if err := tx.Preload("Skin").Where("case_id = ?", parsedCaseID).Find(&contents).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch case contents",
		})
		return
	}

	// select random skin based on drop chances
	selectedContent, err := selectRandomSkin(contents)
	if selectedContent == nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to select skin",
		})
		return
	}


	// Generate random float for skins 
	randomFloat := generateFloat()

	// Calculate value base on the float 
	skin := selectedContent.Skin
	floatMultiplier := 1.0 - randomFloat // better float will result in hight value
	skinValue := skin.MinValue + (skin.MaxValue-skin.MinValue)*floatMultiplier

	// Deduct Case bucks from user 
	balanceBefore := user.Casebucks
	user.Casebucks -= caseItem.Price
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user balance",
		})
		return 
	}

	//
}