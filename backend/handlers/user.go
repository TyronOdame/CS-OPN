package handlers

import (
	"net/http"

	"github.com/TyronOdame/CS-OPN/backend/database"
	"github.com/TyronOdame/CS-OPN/backend/models"
	"github.com/TyronOdame/CS-OPN/backend/middleware"
	"github.com/gin-gonic/gin"
)

// GetProfile returns the current user profile
func GetProfile(c *gin.Context) {
	// takes the information stored form the UserID in the context by the auth middleware
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})
		return 
	}

	var user models.User
	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user.ToJSON(),
	})
}

// updateProfile allows users to update their profile information
type UpdateProfileRequest struct {
	Username string `json:"username" binding:"omitempty,min=3, max=20"`
	Email	string `json:"email" binding:"omitempty,email"`
}

// updateProfile updates the current user's profile
func UpdateProfile(c *gin.Context) {
	// extract user ID from context
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})
		return
	}

	// parse and validate request body
	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	// Get the user from the database
	var user models.User
	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	//update user fields if provide in the request
	if req.Username != "" {
		// check if username is already taken
		var existingUser models.User
		if err := database.DB.Where("username = ? AND id != ?", req.Username, userID).First(&existingUser).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Username already in use",
			})
			return
		}
		user.Username = req.Username
	}

	if req.Email != "" {
		// check if email is already taken
		var existingUser models.User
		if err := database.DB.Where("email = ? AND id != ?", req.Email, userID).First(&existingUser).Error; err == nil{
			c.JSON(http.StatusConflict, gin.H{
				"error": "Email already in use",
			})
			return
		}
		user.Email = req.Email
	}

	// save the updated user to the database
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update profile",
		})
		return 
	}

	// respond with the updated users profile
	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
		"user": user.ToJSON(),

	})
}




