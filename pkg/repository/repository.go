package repository

import (
	"referralmanager-cli-backend/pkg/models"

	"gorm.io/gorm"
	//"referralmanager-cli-backend/pkg/database"
)

// Repository Interface for repositories to provide abstraction
type LeadRepository interface {
	FetchAll() (*gorm.DB, []models.Lead, error)
	FindByCompany(Company string) (*gorm.DB, []models.Lead, error)
	CreateLead(lead *models.Lead) (*gorm.DB, error)
	UpdateLead(lead *models.Lead) (*models.Lead, error)
	DeleteLead(Email string) (string, error)
}
