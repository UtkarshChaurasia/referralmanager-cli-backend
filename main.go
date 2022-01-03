package main

import (
	//"fmt"

	//"referralmanager-cli-backend/pkg/database"
	"referralmanager-cli-backend/pkg/routes"
)

func main() {

	routes.InitializeRouter()

	// db.Create(&Lead{Name: "Test2", Email: "test2@email.com", Message: "Hello Next", Company: "Logmein"})
	// var lead Lead
	// db.First(&lead, "email = ?", "test1@email.com")
	// fmt.Println(lead)
}
