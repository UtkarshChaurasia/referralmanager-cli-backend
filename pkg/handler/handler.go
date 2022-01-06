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
	_, leads, err := l.repo.FetchAll()
	if err != nil {
		errmsg := err.Error()
		json.NewEncoder(res).Encode(errmsg)
	} else {
		json.NewEncoder(res).Encode(leads)
	}

}

// ROUTE 2: Get all the leads of particular company using: GET "/leads/{company}"
func (l *Lead) GetCompanyLeads(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	Company := params["company"]
	result, leads, err := l.repo.FindByCompany(Company)
	rows := result.RowsAffected

	if rows == 0 {
		message := "No Leads of " + Company + " found"
		json.NewEncoder(res).Encode(message)
	}
	if err != nil {
		fmt.Println("Sorry leads for ", Company, " could not be fetched")
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
	result, err := l.repo.CreateLead(newLead)
	rows := result.RowsAffected
	if rows == 0 {
		message := "Another lead with email id: " + newLead.Email + " exists"
		json.NewEncoder(res).Encode(message)
	} else {
		json.NewEncoder(res).Encode(newLead)
	}
	if err != nil {
		fmt.Println("Fetch all leads error")
	} else {
		fmt.Println("Created Lead")
	}

}

// ROUTE 4: Update an existing lead using: PUT "/updatelead"
func (l *Lead) UpdateLead(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var newLead *models.Lead
	json.NewDecoder(req.Body).Decode(&newLead)
	params := mux.Vars(req)

	id := params["id"]
	result, updatedLead, err := l.repo.UpdateLead(newLead, id)
	rows := result.RowsAffected
	if rows == 0 {
		message := "No Leads found with Id " + id + " found"
		json.NewEncoder(res).Encode(message)
	} else {
		json.NewEncoder(res).Encode(updatedLead)
	}
	if err != nil {
		fmt.Println("Update leads error")
	} else {
		fmt.Println("Updated Lead")
	}

}

// ROUTE 5: Delete an existing lead using: DELETE "/updatelead"
func (l *Lead) DeleteLead(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	id := params["id"]
	fmt.Println("id")
	result, err := l.repo.DeleteLead(id)
	rows := result.RowsAffected

	if rows == 0 {
		message := "No Leads found with Id " + id
		json.NewEncoder(res).Encode(message)
	} else {
		message := "Lead with Id " + id + " deleted successfully"
		json.NewEncoder(res).Encode(message)
	}
	if err != nil {
		fmt.Println("Delete lead error")
	} else {
		fmt.Println("Deleted Lead")
	}

}
