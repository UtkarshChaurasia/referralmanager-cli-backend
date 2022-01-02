package routes

import (
	"log"
	"net/http"
	"referralmanager-cli-backend/pkg/repository"

	"github.com/gorilla/mux"
)

// func to initialize router
func InitializeRouter() {

	r := mux.NewRouter()

	r.HandleFunc("/leads", GetAllLeads).Methods("GET")
	r.HandleFunc("/leads/{company}", GetCompanyLeads).Methods("GET")
	r.HandleFunc("/addlead", AddLead).Methods("POST")
	r.HandleFunc("/updatelead", UpdateLead).Methods("PUT")
	r.HandleFunc("/deletelead", DeleteLead).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}

var repo repository.Leadrepository = repository.Newrepository()

// ROUTE 1: Get all the leads using: GET "/leads"
func GetAllLeads(res http.ResponseWriter, req *http.Request) {

}

// ROUTE 2: Get all the leads of particular company using: GET "/leads/{company}"
func GetCompanyLeads(res http.ResponseWriter, req *http.Request) {

}

// ROUTE 3: Add a new lead using: POST "/addlead"
func AddLead(res http.ResponseWriter, req *http.Request) {

}

// ROUTE 4: Update an existing lead using: PUT "/updatelead"
func UpdateLead(res http.ResponseWriter, req *http.Request) {

}

// ROUTE 5: Delete an existing lead using: DELETE "/updatelead"
func DeleteLead(res http.ResponseWriter, req *http.Request) {

}
