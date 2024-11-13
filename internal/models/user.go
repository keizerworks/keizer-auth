package models

import "time"

type User struct {
	LastLogin    time.Time `json:"last_login"`
	Email        string    `gorm:"not null;default:null;unique,index;type:varchar(100)" json:"email"`
	PasswordHash string    `json:"-"`
	FirstName    string    `gorm:"not null;type:varchar(100);default:null" json:"first_name"`
	LastName     string    `gorm:"type:varchar(100);default:null" json:"last_name"`
	Base
	IsVerified bool `gorm:"not null;default:false" json:"is_verified"`
	IsActive   bool `gorm:"not null;default:false" json:"is_active"`
}
