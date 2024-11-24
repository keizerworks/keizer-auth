package services

import (
	"keizer-auth/internal/models"
	"keizer-auth/internal/repositories"

	"github.com/google/uuid"
)

type AccountService struct {
	accountRepo     *repositories.AccountRepository
	userAccountRepo *repositories.UserAccountRepository
}

func NewAccountService(
	accountRepo *repositories.AccountRepository,
	userAccountRepo *repositories.UserAccountRepository,
) *AccountService {
	return &AccountService{accountRepo: accountRepo, userAccountRepo: userAccountRepo}
}

func (self *AccountService) Create(
	name string,
	userID uuid.UUID,
) (*models.Account, error) {
	account := models.Account{Name: name}
	if err := self.accountRepo.Create(&account); err != nil {
		return nil, err
	}

	userAccount := models.UserAccount{
		UserID:    userID,
		AccountID: account.ID,
		Role:      "admin",
	}

	if err := self.userAccountRepo.Create(&userAccount); err != nil {
		return nil, err
	}

	return &account, nil
}

func (self *AccountService) GetAccountsByUser(
	userID uuid.UUID,
) (*[]models.Account, error) {
	accounts, err := self.accountRepo.GetAccountsByUser(userID)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
