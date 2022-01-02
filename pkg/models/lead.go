package models

type Lead struct {
	Name    string
	Email   string `gorm:"primaryKey"`
	Message string
	Company string
}
