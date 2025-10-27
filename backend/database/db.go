package database

import (
	"fmt"
	"log"

	"github.com/TyronOdame/CS-OPN/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB initializes the database connection using GORM
var DB *gorm.DB

// ConnectDB establishes a connection to the PostgreSQL database
func ConnectDB(host, port, user, password, dbname string) (*gorm.DB, error) {
	// build Postgres connection string
	dsn := fmt.Sprint(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname,
	)

	// open a connection to the database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})	

	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}

	// Auto-migrate the models to create/update database tables
	err = DB.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		log.Printf("Failed to auto-migrate database: %v", err)
		return nil, err
	}

	log.Println("Database connection established and models migrated")
	return DB, nil
		

}
