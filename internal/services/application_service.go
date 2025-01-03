package services

import (
	"errors"
	"keizer-auth/internal/models"
	"keizer-auth/internal/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ApplicationService struct {
	applicationRepo *repositories.ApplicationRepository
	accountRepo     *repositories.AccountRepository
}

func NewApplicationService(
	applicationRepo *repositories.ApplicationRepository,
	accountRepo *repositories.AccountRepository,
) *ApplicationService {
	return &ApplicationService{
		applicationRepo: applicationRepo,
		accountRepo:     accountRepo,
	}
}

func (self *ApplicationService) Create(
	name string,
	accountID uuid.UUID,
	userID uuid.UUID,
) (*models.Application, error) {
	_, err := self.accountRepo.GetAccountByUser(accountID, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("account not found or unauthorized access")
		}
		return nil, err
	}

	application := models.Application{Name: name, AccountID: accountID}
	if err := self.applicationRepo.Create(&application); err != nil {
		return nil, err
	}
	return &application, nil
}

func (self *ApplicationService) Get(
	accountID uuid.UUID,
	userID uuid.UUID,
) (*[]models.Application, error) {
	_, err := self.accountRepo.GetAccountByUser(accountID, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("account not found or unauthorized access")
		}
		return nil, err
	}

	applications, err := self.applicationRepo.GetApplicationsByAccount(accountID)
	if err != nil {
		return nil, err
	}

	return applications, nil
}
