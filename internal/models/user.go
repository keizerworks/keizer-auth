package models

import "time"

type User struct {
	LastLogin    time.Time
	Email        string `gorm:"not null;unique,index;type:varchar(100)"`
	PasswordHash string
	FirstName    string `gorm:"not null;type:varchar(100)"`
	LastName     string `gorm:"type:varchar(100)"`
	Base
	Sessions   []Session
	IsVerified bool `gorm:"not null;default:false"`
	IsActive   bool `gorm:"not null;default:false"`
	Otp        string
}
