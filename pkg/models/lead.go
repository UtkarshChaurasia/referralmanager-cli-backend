package models

import (
	"time"

	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Name      string
	Email     string `gorm:"unique"`
	Message   string
	Company   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
