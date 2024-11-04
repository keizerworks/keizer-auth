package models

import (
	"time"

	"github.com/google/uuid"
)

// Base contains common columns for all tables.
type Base struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid()"`
}
