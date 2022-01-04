package database

import (
	"fmt"
	"referralmanager-cli-backend/pkg/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//var DBConn *gorm.DB

type DB struct {
	SQL *gorm.DB
}

var dbConn = &DB{}

// function to connect with database
func ConnectDatabase() (*DB, error) {
	db, err := gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to connect database")
	}
	//sqlDB, _ := db.DB()
	//defer sqlDB.Close()
	db.AutoMigrate(&models.Lead{})
	dbConn.SQL = db
	return dbConn, err
}
