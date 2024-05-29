package database

import (
	//models

	"Basket/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Panicf("Failed to connect to database: %v\n", err)
	}

	log.Println("Database connection established")
	err = DB.AutoMigrate(&model.User{})
	err = DB.AutoMigrate(&model.Basket{})

	if err != nil {
		log.Panicf("Failed to migrate the database: %v\n", err)
	}
}
