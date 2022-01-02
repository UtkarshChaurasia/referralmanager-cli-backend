package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

// function to connect with database
func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
