package models

import (
	"fmt"
	"github.com/henvo/golang-gin-gorm-starter/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is the he database connection.
var DB *gorm.DB

// SetupDatabase migrates and sets up the database.
func SetupDatabase() {
	u := helper.GetEnv("DATABASE_USER", "golang")
	p := helper.GetEnv("DATABSE_PASSWORD", "golang")
	h := helper.GetEnv("DATABASE_HOST", "localhost:3306")
	n := helper.GetEnv("DATABASE_NAME", "go_test")
	q := "charset=utf8mb4&parseTime=True&loc=Local"

	// Assemble the connection string.
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", u, p, h, n, q)

	// Connect to the database.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Migrate the schema
	db.AutoMigrate(&User{})

	if err != nil {
		panic("Could not open database connection")
	}

	DB = db
}
