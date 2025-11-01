package main

import (
	"log"

	"github.com/TyronOdame/CS-OPN/backend/database"
	"github.com/TyronOdame/CS-OPN/backend/handlers"
	"github.com/gin-gonic/gin"
)

// starting the main function to run backend and database connection
func main() {
	// loading configuration for .env file
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal(" Failed to load config:", err)
	}

	// connect to the database (PostgreSQL)
	err = database.ConnectDB(
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
	if err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}

	// Run the migrations
	err = database.AutoMigrate()
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	// Create HTTP server 
	router := gin.Default()

	// Health Check Endpoints
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":   "ok",
			"database": "connected",
			"message":  "CS:OPN backend is running!",
		})
	})

	//Auth routes
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", handlers.RegisterHandler(cfg.JWTSecret))
		authRoutes.POST("/login", handlers.Login(cfg.JWTSecret))
	}

	// Start server
	log.Printf("🚀 Server starting on port %s", cfg.ServerPort)
	log.Printf("📍 Health check: http://localhost:%s/health", cfg.ServerPort)
	router.Run(":" + cfg.ServerPort)


	
}