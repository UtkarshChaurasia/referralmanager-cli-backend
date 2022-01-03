package handler

import (
	"net/http"
	"referralmanager-cli-backend/pkg/database"
	"referralmanager-cli-backend/pkg/repository"
	"referralmanager-cli-backend/pkg/repository/lead_sqlite"
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

}

// ROUTE 2: Get all the leads of particular company using: GET "/leads/{company}"
func (l *Lead) GetCompanyLeads(res http.ResponseWriter, req *http.Request) {

}

// ROUTE 3: Add a new lead using: POST "/addlead"
func (l *Lead) AddLead(res http.ResponseWriter, req *http.Request) {

}

// ROUTE 4: Update an existing lead using: PUT "/updatelead"
func (l *Lead) UpdateLead(res http.ResponseWriter, req *http.Request) {

}

// ROUTE 5: Delete an existing lead using: DELETE "/updatelead"
func (l *Lead) DeleteLead(res http.ResponseWriter, req *http.Request) {

}
