package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ApplicationEnvironment struct {
	Name   string `gorm:"not null;type:varchar(50)" json:"name"`
	Status string `gorm:"type:varchar(50);default:'active'" json:"status"`
	Base
	Application   Application `gorm:"foreignKey:ApplicationID"`
	ApplicationID uuid.UUID   `gorm:"type:uuid;not null;index" json:"application_id"`
	IsProtected   bool        `gorm:"not null;default:false" json:"is_protected"`
}

func (e *ApplicationEnvironment) BeforeCreate(tx *gorm.DB) error {
	if e.Name == "development" || e.Name == "production" {
		e.IsProtected = true
	}
	return nil
}
