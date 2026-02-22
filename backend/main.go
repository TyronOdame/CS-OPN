package main

import (
	"log"

	"github.com/TyronOdame/CS-OPN/backend/database"
	"github.com/TyronOdame/CS-OPN/backend/handlers"
	"github.com/TyronOdame/CS-OPN/backend/middleware"
	"github.com/TyronOdame/CS-OPN/backend/seed"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// starting the main function to run backend and database connection
func main() {
	// loading configuration for .env file
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal("❌ Failed to load config:", err)
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

	if err := database.SeedDatabase(); err != nil {
		log.Fatal("❌ Database seeding failed:", err)
	}

	//  Seed test data
	log.Println("🌱 Starting database seeding...")
	seed.SeedCases()
	seed.SeedSkins()        
	seed.SeedCaseSkins()
	log.Println("✅ Database seeding complete!")

	// Create HTTP server 
	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * 3600, // 12 hours
	}))

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

	// User routes (protected - requires JWT token)
	userRoutes := router.Group("/user")
	userRoutes.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		userRoutes.GET("/profile", handlers.GetProfile)
		userRoutes.PUT("/profile", handlers.UpdateProfile)
	}

	// Case routes
	caseRoutes := router.Group("/cases")
	{
		// public routes
		caseRoutes.GET("", handlers.GetAllCases)
		caseRoutes.GET("/:id", handlers.GetCaseByID)

		// protected routes
		caseRoutes.POST("/:id/buy", middleware.AuthMiddleware(cfg.JWTSecret), handlers.BuyCase)
		caseRoutes.POST("/:id/open", middleware.AuthMiddleware(cfg.JWTSecret), handlers.OpenCase)
	}

	// Inventory routes (protected)
	inventoryRoutes := router.Group("/inventory")
	inventoryRoutes.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		inventoryRoutes.GET("", handlers.GetUserInventory)
		inventoryRoutes.POST("/:id/sell", handlers.SellInventoryItem)
		inventoryRoutes.GET("/cases", handlers.GetUserCases)
		inventoryRoutes.POST("/cases/:id/open", handlers.OpenPurchasedCase)
	}

	// Transaction routes (protected)
	transactionRoutes := router.Group("/transactions")
	transactionRoutes.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		transactionRoutes.GET("", handlers.GetUserTransactions)
	
	}

	// AI price check (mocked provider for v1)
	aiRoutes := router.Group("/ai")
	aiRoutes.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		aiRoutes.POST("/price-check", handlers.PriceCheckMock)
	}

	// Start server
	log.Printf("🚀 Server starting on port %s", cfg.ServerPort)
	log.Printf("📍 Health check: http://localhost:%s/health", cfg.ServerPort)
	log.Printf("🎁 Cases: GET http://localhost:%s/cases", cfg.ServerPort)
	log.Printf("👤 Profile: GET http://localhost:%s/user/profile (protected)", cfg.ServerPort)
	log.Printf("✏️  Update: PUT http://localhost:%s/user/profile (protected)", cfg.ServerPort)
	log.Printf("🔐 Register: POST http://localhost:%s/auth/register", cfg.ServerPort)
	router.Run(":" + cfg.ServerPort)


	
}