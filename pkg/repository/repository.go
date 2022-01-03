package repository

import (
	"referralmanager-cli-backend/pkg/models"
	//"referralmanager-cli-backend/pkg/database"
)

// Repository Interface for repositories to provide abstraction
type LeadRepository interface {
	FetchAll() ([]models.Lead, error)
	FindByCompany(Company string) ([]models.Lead, error)
	CreateLead(lead *models.Lead) (*models.Lead, error)
	UpdateLead(lead *models.Lead) (*models.Lead, error)
	DeleteLead(Email string) (string, error)
}
