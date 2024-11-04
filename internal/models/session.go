package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ExpiresAt time.Time
	Token     string `gorm:"not null;unique"`
	SessionId string `gorm:"primaryKey;not null"`
	UserID    uuid.UUID
}
