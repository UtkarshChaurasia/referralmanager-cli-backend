package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRouter() {

	r := mux.NewRouter()
	r.HandleFunc("/leads", GetAllLeads).Methods("GET")
	r.HandleFunc("/leads/{company}", GetCompanyLeads).Methods("GET")
	r.HandleFunc("/addlead", AddLead).Methods("POST")
	r.HandleFunc("/updatelead/{email}", UpdateLead).Methods("PUT")
	r.HandleFunc("/deletelead/{email}", DeleteLead).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}
