package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB is the he database connection.
var DB *gorm.DB

// SetupDatabase migrates and sets up the database.
func SetupDatabase() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Could not open database connection")
	}

	DB = db
}
