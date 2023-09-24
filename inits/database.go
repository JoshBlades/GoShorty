package inits

import (
	"log"

	"github.com/joshblades/goshorty/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("goshorty.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database successfully")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")

	// TODO: Add migreations
	db.AutoMigrate(&models.Link{}, &models.Click{})

	Database = db

}
