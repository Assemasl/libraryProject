package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"book-library/config"
	"book-library/models"
	"book-library/routes"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := config.GetDatabaseDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Run migrations manually or via goose separately
	db.AutoMigrate(&models.User{}, &models.Book{}, &models.Reaction{})

	app := fiber.New()

	// Set DB to global
	models.SetDB(db)

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
