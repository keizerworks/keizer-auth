package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base contains common columns for all tables.
type Base struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	ID        uuid.UUID  `json:"id" gorm:"type:uuid"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID, err = uuid.NewV7()
	return err
}
