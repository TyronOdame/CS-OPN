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

	// get all skins in this casa with drop chances
	var contents []models.CaseContent
	if err := database.DB.Preload("Skin").Where("case_id = ?", caseID).Find(&contents).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Case not found",
			
		})
		return
	}
	

	// 
}