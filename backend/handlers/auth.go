package handlers

import (
	"net/http"

	"github.com/TyronOdame/CSC-Project/backend/database"
	"github.com/TyronOdame/CSC-Project/backend/utils"
	"github.com/TyronOdame/CSC-Project/backend/models"
	"github.com/gin-gonic/gin"
)

// RegisterRequest represents the expected payload for user registration
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginRequest represents the expected payload for user login
type LoginRequest struct {
	Email   string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// This is the register handler
func RegisterHandler(jwtSecret string) gin.HandlerFunc {
	return func (c *gin.Context) {
		var req RegisterRequest

		// validate request body
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request data",
				"details": err.Error(),
			})
			return
		}

		// check if email already exists
		var existingUser models.User
		if err := database.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Email already in use",
			})
			return
		}

		//check if username already exists
		if err := database.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Username already in use",
			})
			return
		}
			

		// create new user
		user := models.User{
			Email: req.Email,
			Username: req.Username,
		}

		// hashing the password
		if err := user.HashPassword(req.Password); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to hash password",
			})
			return
	
		}
	}
}
		
	

			
