package repositories

import (
	"fmt"
	"keizer-auth/internal/models"

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

func (r *UserRepository) GetUser(uuid string) (*models.User, error) {
	user := new(models.User)
	result := r.db.First(&user, "id = ?", uuid)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error in getting user: %w", result.Error)
	}

	return user, nil
}
