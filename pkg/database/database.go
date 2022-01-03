package database

import (
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
		panic("failed to connect database")
	}

	dbConn.SQL = db
	return dbConn, err
}
