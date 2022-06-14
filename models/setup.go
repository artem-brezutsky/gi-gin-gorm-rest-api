package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/golang"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	database.AutoMigrate(&Book{})
	DB = database
}
