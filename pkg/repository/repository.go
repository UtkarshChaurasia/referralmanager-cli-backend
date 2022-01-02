package repository

import "referralmanager-cli-backend/pkg/models"

// Interface for repositorys to provide abstraction
type Leadrepository interface {
	FetchAll() ([]models.Lead, error)
	FindByCompany(Company string) ([]models.Lead, error)
	CreateLead(lead *models.Lead) (*models.Lead, error)
	UpdateLead(lead *models.Lead) (*models.Lead, error)
	DeleteLead(Email string) (string, error)
}

type repository struct{}

func Newrepository() Leadrepository {
	return &repository{}
}

// function to fetch all the leads
func (*repository) FetchAll() ([]models.Lead, error) {
	var leads []models.Lead

	return leads, nil
}

// function to fetch all the leads of particular company
func (*repository) FindByCompany(Company string) ([]models.Lead, error) {
	var leads []models.Lead

	return leads, nil
}

// function to create a new lead
func (*repository) CreateLead(lead *models.Lead) (*models.Lead, error) {
	return lead, nil
}

// function to Update an existing lead
func (*repository) UpdateLead(lead *models.Lead) (*models.Lead, error) {
	return lead, nil
}

// function to delete an existing lead
func (*repository) DeleteLead(Email string) (string, error) {
	return "", nil
}
