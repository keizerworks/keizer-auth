package repositories

import (
	"fmt"
	"keizer-auth/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (self *AccountRepository) Create(account *models.Account) error {
	return self.
		db.Model(account).
		Clauses(clause.Returning{}).
		Create(account).
		Error
}

func (self *AccountRepository) GetAccountsByUser(
	userID uuid.UUID,
) (*[]models.Account, error) {
	accounts := new([]models.Account)
	if err := self.db.
		Joins("JOIN user_accounts ON user_accounts.account_id = accounts.id").
		Where("user_accounts.user_id = ?", userID).
		Find(accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func (self *AccountRepository) GetAccountByUser(
	accountID uuid.UUID,
	userID uuid.UUID,
) (*models.Account, error) {
	account := new(models.Account)
	if err := self.
		db.Model(models.Account{}).
		Joins("JOIN user_accounts ON user_accounts.account_id = accounts.id").
		Where("user_accounts.user_id = ? AND accounts.id = ?", userID, accountID).
		First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func (self *AccountRepository) Get(uuid string) (*models.Account, error) {
	account := new(models.Account)
	result := self.db.First(&account, "id = ?", uuid)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("account not found")
		}
		return nil, fmt.Errorf("error in getting account: %w", result.Error)
	}

	return account, nil
}
