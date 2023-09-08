package database

import (
	"github.com/wellrcosta/Api.Crud.Go.Gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {

	connectionString := "host=192.168.0.115 user=root password=root dbname=projects port=9432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL database: %v", err)
	}
	// Auto-migrate your models after successful connection
	DB.AutoMigrate(&models.Student{})
}
