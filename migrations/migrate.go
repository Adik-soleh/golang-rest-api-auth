package migrations

import (
	"golang-rest-api/config"
	"golang-rest-api/models"
	"log"
)

func RunMigration() {
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Migration completed successfully")
}
