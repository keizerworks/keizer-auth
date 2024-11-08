package models

import "time"

type User struct {
	LastLogin    time.Time `json:"lastLogin"`
	Email        string    `gorm:"not null;default:null;unique,index;type:varchar(100)" json:"email"`
	PasswordHash string    `json:"-"`
	FirstName    string    `gorm:"not null;type:varchar(100);default:null" json:"fName"`
	LastName     string    `gorm:"type:varchar(100);default:null" json:"lName"`
	Base
	Sessions   []Session
	IsVerified bool `gorm:"not null;default:false" json:"isVerified"`
	IsActive   bool `gorm:"not null;default:false" json:"isActive"`
}
