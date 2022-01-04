package models

import (
	"time"

	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name      string
	Email     string
	Message   string `gorm:"primaryKey"`
	Company   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
