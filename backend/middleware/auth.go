package middleware

import (
	"net/http"
	"strings"

	"github.com/TyronOdame/CS-OPN/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Function to create authentication middleware
func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// here is where the middleware logic will go

		// get Auth header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required",
			})
			c.Abort() // this will kill the processing of this request
			return
		}

		// I'm Expecting format: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization header format. Use: Bearer <token>",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// validating the JWT token for Authentication
		claims, err := utils.ValidateJWT(tokenString, jwtSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// token is valid! store user info in context for handlers to use
		c.Set("userID", claims.UserID)
		c.Set("userEmail", claims.Email)
		c.Set("username", claims.Username)

		// continue to the next handler
		c.Next()

	}
}