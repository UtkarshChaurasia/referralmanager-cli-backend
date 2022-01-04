package lead_sqlite

import (
	"fmt"
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

// function to fetch all the leads
func (m *sqliteLeadRepo) FetchAll() (*gorm.DB, []models.Lead, error) {
	var leads []models.Lead

	result := m.Conn.Find(&leads)
	err := result.Error
	return result, leads, err
}

// function to fetch all the leads of particular company
func (m *sqliteLeadRepo) FindByCompany(Company string) (*gorm.DB, []models.Lead, error) {
	var leads []models.Lead

	result := m.Conn.Where("Company = ?", Company).Find(&leads)
	err := result.Error
	return result, leads, err
}

// function to create a new lead
func (m *sqliteLeadRepo) CreateLead(lead *models.Lead) (*gorm.DB, error) {
	result := m.Conn.Create(&lead)
	err := result.Error
	return result, err
}

// function to Update an existing lead
func (m *sqliteLeadRepo) UpdateLead(lead *models.Lead, ID string) (*gorm.DB, *models.Lead, error) {
	var newLead models.Lead
	fmt.Println(ID)
	m.Conn.Where("ID = ?", ID).Find(&newLead)
	fmt.Println(newLead)
	newLead.Name = lead.Name
	newLead.Email = lead.Email
	newLead.Message = lead.Message
	newLead.Company = lead.Company
	result := m.Conn.Save(&newLead)
	err := result.Error
	if err != nil {
		fmt.Println("Update lead error")
	} else {
		fmt.Println("Updated Successfully")
	}
	return result, &newLead, err
}

// function to delete an existing lead
func (m *sqliteLeadRepo) DeleteLead(ID string) (*gorm.DB, error) {
	var lead models.Lead

	fmt.Println(ID)
	result := m.Conn.Where("ID = ?", ID).Delete(&lead)
	err := result.Error
	if err != nil {
		fmt.Println("Delete lead error")
	} else {
		fmt.Println("Deleted Successfully")
	}
	return result, err
}
