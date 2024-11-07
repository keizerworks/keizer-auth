package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ExpiresAt time.Time
	Token     string `gorm:"not null;unique;default:null"`
	SessionId string `gorm:"primaryKey;not null;default:null"`
	UserID    uuid.UUID
}
