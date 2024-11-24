package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserType string

const (
	Dashboard UserType = "dashboard"
	Member    UserType = "member"
)

func (self *UserType) Scan(value interface{}) error {
	if value == "" {
		*self = Member
		return nil
	}

	strVal, ok := value.(string)
	if !ok {
		return fmt.Errorf("Failed to convert")
	}

	*self = UserType(strVal)
	return nil
}

func (self UserType) Value() (driver.Value, error) {
	return string(self), nil
}

func (ut UserType) Validate() bool {
	switch ut {
	case Dashboard, Member:
		return true
	default:
		return false
	}
}

type User struct {
	LastLogin    time.Time `json:"last_login"`
	Email        string    `gorm:"not null;default:null;unique,index;type:varchar(100)" json:"email"`
	PasswordHash string    `json:"-"`
	FirstName    string    `gorm:"not null;type:varchar(100);default:null" json:"first_name"`
	LastName     string    `gorm:"type:varchar(100);default:null" json:"last_name"`
	Type         UserType  `gorm:"type:user_type;not null;default:'member'" json:"type"`
	Base
	IsVerified bool `gorm:"not null;default:false" json:"is_verified"`
	IsActive   bool `gorm:"not null;default:false" json:"is_active"`
}

func (u *User) BeforeMigrate(db *gorm.DB) error {
	return db.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_type') THEN
				CREATE TYPE user_type AS ENUM ('dashboard', 'member');
			END IF;
		END$$;
	`).Error
}
