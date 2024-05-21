package bootstrap

import (
	"fmt"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"project_weather/config"
	"project_weather/models"
)

var FXModule_Persistence = fx.Module(
	"persistence",
	fx.Provide(
		createDatabaseConnection,
	),

	fx.Invoke(
		createDatabaseSchema,
	),
)

func createDatabaseConnection(config *config.ApplicationConfiguration) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%v port=%v user=%v dbname=%v password=%v",
		config.DBHost, config.DBPort, config.DBUsername, config.DBName, config.DBPassword,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func createDatabaseSchema(db *gorm.DB) error {

	err := db.AutoMigrate(&models.City{}, &models.Forecast{})
	if err != nil {
		log.Fatal("Failed to migrate database schema:", err)
		return err
	}
	return nil
}
