package routes

import (
	"log"
	"net/http"
	"referralmanager-cli-backend/pkg/database"
	"referralmanager-cli-backend/pkg/handler"

	"github.com/gorilla/mux"
)

// func to initialize router
func InitializeRouter() {

	Connection, err := database.ConnectDatabase()
	if err != nil {
		panic("Database error")
	}

	lHandler := handler.NewLeadHandler(Connection)

	r := mux.NewRouter()

	r.HandleFunc("/leads", lHandler.GetAllLeads).Methods("GET")
	r.HandleFunc("/leads/{company}", lHandler.GetCompanyLeads).Methods("GET")
	r.HandleFunc("/addlead", lHandler.AddLead).Methods("POST")
	r.HandleFunc("/updatelead", lHandler.UpdateLead).Methods("PUT")
	r.HandleFunc("/deletelead", lHandler.UpdateLead).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}

//var repo repository.Leadrepository = repository.Newrepository()
