package database

import (
	"fmt"
	"log"

	"github.com/TyronOdame/CS-OPN/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB initializes the database connection using GORM
var DB *gorm.DB

// ConnectDB establishes a connection to the PostgreSQL database
func ConnectDB(host, port, user, password, dbname string)  error {
	// build Postgres connection string
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname,
	)

	// open a connection to the database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})	

	if err != nil {
		return fmt.Errorf("Failed to connect to database: %w", err)
	}

	log.Println("Database connection successful")
	return nil
}



// Auto-migrate the models to create/update database tables
func AutoMigrate() error {
	log.Println("Running database migrations...")
	err := DB.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		return fmt.Errorf("AutoMigrate failed: %w", err)
	}

	log.Println("Database migrations completed successfully")
	return nil
}
	
