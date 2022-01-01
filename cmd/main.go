package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Lead struct {
	Name    string
	Email   string
	Message string
	Company string
}

func main() {
	db, err := gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Create(&Lead{Name: "Test2", Email: "test2@email.com", Message: "Hello Next", Company: "Logmein"})
	var lead Lead
	db.First(&lead, "email = ?", "test1@email.com")
	fmt.Println(lead)
}
