package repositories

import (
	"fmt"
	"keizer-auth/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ApplicationRepository struct {
	db *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) *ApplicationRepository {
	return &ApplicationRepository{db: db}
}

func (self *ApplicationRepository) Create(application *models.Application) error {
	return self.
		db.Model(application).
		Clauses(clause.Returning{}).
		Create(application).
		Error
}

func (self *ApplicationRepository) GetByID(uuid string) (*models.Application, error) {
	application := new(models.Application)
	result := self.db.First(&application, "id = ?", uuid)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("application not found")
		}
		return nil, fmt.Errorf("error in getting application: %w", result.Error)
	}

	return application, nil
}

func (self *ApplicationRepository) GetApplicationsByAccount(
	accountID uuid.UUID,
) (*[]models.Application, error) {
	applications := new([]models.Application)
	err := self.db.
		Where("account_id = ?", accountID).
		Find(applications).Error
	if err != nil {
		return nil, err
	}
	return applications, nil
}
