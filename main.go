package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lpernett/godotenv"
	"log"
	"os"
	"project_weather/models"
	"project_weather/repositories"
	"project_weather/storage"
)

func main() {
	err := godotenv.Load("resources/environments/local.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DB"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
	}
	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal(err)
	}
	err = models.MigrateCity(db)
	if err != nil {
		log.Fatal(err)
	}

	repo := &repositories.CityRepository{
		DB: db,
	}
	app := fiber.New()
	repo.SetupRoutes(app)
	app.Listen(":8080")
}
