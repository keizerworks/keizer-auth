package models

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserAccountRole string

const (
	RoleAdmin  UserAccountRole = "admin"
	RoleMember UserAccountRole = "member"
)

type Account struct {
	Name string `gorm:"not null;default:null;type:varchar(100)" json:"name"`
	Logo string `gorm:"default:null" json:"logo"`
	Base
	Users []User `gorm:"many2many:user_accounts" json:"-"`
}

type UserAccount struct {
	UniqueConstraint string          `gorm:"uniqueIndex:user_account_unique,priority:1" json:"-"`
	Role             UserAccountRole `gorm:"not null;type:varchar(50);default:'member'"`
	Account          Account         `gorm:"foreignKey:AccountID"`
	Base
	User      User      `gorm:"foreignKey:UserID"`
	AccountID uuid.UUID `gorm:"type:uuid;not null;index" json:"account_id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
}

func (self UserAccountRole) IsValid() bool {
	switch self {
	case RoleAdmin, RoleMember:
		return true
	default:
		return false
	}
}

func (self *UserAccount) BeforeSave(tx *gorm.DB) error {
	if !self.Role.IsValid() {
		return fmt.Errorf("invalid role: %s", self.Role)
	}
	return nil
}
