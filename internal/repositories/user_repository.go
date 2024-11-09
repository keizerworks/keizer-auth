package repositories

import (
	"fmt"
	"keizer-auth-api/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetUser(uuid uuid.UUID) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, uuid.String())
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error in getting user: %w", result.Error)
	}

	return &user, nil
}
