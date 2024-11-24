package models

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Application struct {
	Name string `gorm:"not null;default:null;type:varchar(100)" json:"name"`
	Logo string `gorm:"default:null" json:"logo"`
	Base
	Account   Account   `gorm:"foreignKey:AccountID"`
	AccountID uuid.UUID `gorm:"type:uuid;not null;index" json:"account_id"`
}

func (a *Application) AfterCreate(tx *gorm.DB) error {
	environments := []ApplicationEnvironment{
		{Name: "development", ApplicationID: a.ID},
		{Name: "production", ApplicationID: a.ID},
	}
	if err := tx.Create(&environments).Error; err != nil {
		return errors.New("Failed to create default environments")
	}
	return nil
}
