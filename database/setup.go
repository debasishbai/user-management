package database

import (
	"user-management/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.User{})
	database.LogMode(true)
	//create dummy data
	// database.Create(&models.User{FirstName: "Bruce", LastName: "Wayne", Age: 30, Occupation: "Engineer"})
	DB = database
}
