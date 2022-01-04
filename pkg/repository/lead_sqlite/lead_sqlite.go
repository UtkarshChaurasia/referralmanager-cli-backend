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

// func Newrepository() LeadRepository {
// 	return &repository{}
// }

// function to fetch all the leads
func (m *sqliteLeadRepo) FetchAll() (*gorm.DB, []models.Lead, error) {
	var leads []models.Lead

	result := m.Conn.Find(&leads)
	//result := m.Conn.Exec("SELECT * FROM leads")
	//sqlDB, _ := m.Conn.DB()
	//sqlDB.Close()
	//m.Conn.Find(&leads)
	return result, leads, nil
}

// function to fetch all the leads of particular company
func (m *sqliteLeadRepo) FindByCompany(Company string) (*gorm.DB, []models.Lead, error) {
	var leads []models.Lead

	result := m.Conn.Where("Company = ?", Company).Find(&leads)
	//sqlDB, _ := m.Conn.DB()
	//sqlDB.Close()
	return result, leads, nil
}

// function to create a new lead
func (m *sqliteLeadRepo) CreateLead(lead *models.Lead) (*gorm.DB, error) {

	// dummy := []models.Lead{
	// 	{Name: "Test1", Email: "test1@email.com", Message: "Hello", Company: "RedHat"},
	// }
	// for _, u := range dummy {
	// 	m.Conn.Create(&u)
	// }
	createdLead := m.Conn.Create(&lead)

	//sqlDB, _ := m.Conn.DB()
	//sqlDB.Close()
	return createdLead, nil
}

// function to Update an existing lead
func (m *sqliteLeadRepo) UpdateLead(lead *models.Lead) (*models.Lead, error) {
	return lead, nil
}

// function to delete an existing lead
func (m *sqliteLeadRepo) DeleteLead(ID string) (string, error) {
	var lead models.Lead
	var message string
	fmt.Println(ID)
	result := m.Conn.Where("ID = ?", ID).Delete(&lead)
	err := result.Error
	if err != nil {
		fmt.Println("Delete lead error")
	} else {
		fmt.Println("Deleted Successfully")
	}
	return message, err
}
