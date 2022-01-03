package lead_sqlite

import (
	"referralmanager-cli-backend/pkg/models"
	"referralmanager-cli-backend/pkg/repository"

	"gorm.io/gorm"
)

func NewSQLLeadRepo(Conn *gorm.DB) repository.LeadRepository {
	return &sqliteLeadRepo{
		Conn: Conn,
	}
}

type sqliteLeadRepo struct {
	Conn *gorm.DB
}

// func Newrepository() LeadRepository {
// 	return &repository{}
// }

// function to fetch all the leads
func (m *sqliteLeadRepo) FetchAll() ([]models.Lead, error) {
	var leads []models.Lead

	return leads, nil
}

// function to fetch all the leads of particular company
func (m *sqliteLeadRepo) FindByCompany(Company string) ([]models.Lead, error) {
	var leads []models.Lead

	return leads, nil
}

// function to create a new lead
func (m *sqliteLeadRepo) CreateLead(lead *models.Lead) (*models.Lead, error) {
	return lead, nil
}

// function to Update an existing lead
func (m *sqliteLeadRepo) UpdateLead(lead *models.Lead) (*models.Lead, error) {
	return lead, nil
}

// function to delete an existing lead
func (m *sqliteLeadRepo) DeleteLead(Email string) (string, error) {
	return "", nil
}