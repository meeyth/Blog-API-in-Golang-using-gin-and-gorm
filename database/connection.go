package database

import (
	"social-media/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Request body for user
type ReqBody struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Birthday string `json:"birthday"`
}

var db *gorm.DB

func ConnectToDb() {
	db_conn, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db = db_conn

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
}
