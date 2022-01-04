package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"referralmanager-cli-backend/pkg/database"
	"referralmanager-cli-backend/pkg/models"
	"referralmanager-cli-backend/pkg/repository"
	"referralmanager-cli-backend/pkg/repository/lead_sqlite"

	"github.com/gorilla/mux"
)

func NewLeadHandler(db *database.DB) *Lead {
	return &Lead{
		repo: lead_sqlite.NewSQLLeadRepo(db.SQL),
	}
}

type Lead struct {
	repo repository.LeadRepository
}

// ROUTE 1: Get all the leads using: GET "/leads"
func (l *Lead) GetAllLeads(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	//var allleads []models.Lead
	_, leads, err := l.repo.FetchAll()
	if err != nil {
		fmt.Println("Fetch all leads error")
	} else {
		fmt.Println("Fetched Successfully")
	}
	//fmt.Println("{}", leads)
	json.NewEncoder(res).Encode(leads)

}

// ROUTE 2: Get all the leads of particular company using: GET "/leads/{company}"
func (l *Lead) GetCompanyLeads(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	Company := params["company"]
	fmt.Println(Company)
	_, leads, err := l.repo.FindByCompany(Company)

	if err != nil {
		fmt.Println("Fetch all leads error")
	} else {
		fmt.Println("Fetched Successfully")
	}

	fmt.Println("Company Route")
	json.NewEncoder(res).Encode(leads)

}

// ROUTE 3: Add a new lead using: POST "/addlead"
func (l *Lead) AddLead(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	var newLead *models.Lead
	json.NewDecoder(req.Body).Decode(&newLead)
	_, err := l.repo.CreateLead(newLead)
	if err != nil {
		fmt.Println("Fetch all leads error")
	} else {
		fmt.Println("Created Lead")
	}
	json.NewEncoder(res).Encode(newLead)

}

// ROUTE 4: Update an existing lead using: PUT "/updatelead"
func (l *Lead) UpdateLead(res http.ResponseWriter, req *http.Request) {

}

// ROUTE 5: Delete an existing lead using: DELETE "/updatelead"
func (l *Lead) DeleteLead(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	id := params["id"]
	fmt.Println("id")
	message, err := l.repo.DeleteLead(id)
	if err != nil {
		fmt.Println("Delete lead error")
	} else {
		fmt.Println("Deleted Lead")
	}
	json.NewEncoder(res).Encode(message)

}
