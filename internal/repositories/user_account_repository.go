package repositories

import (
	"keizer-auth/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserAccountRepository struct {
	db *gorm.DB
}

func NewUserAccountRepository(db *gorm.DB) *UserAccountRepository {
	return &UserAccountRepository{db: db}
}

func (self *UserAccountRepository) Create(
	userAccount *models.UserAccount,
) error {
	return self.
		db.Model(userAccount).
		Clauses(clause.Returning{}).
		Create(userAccount).
		Error
}
