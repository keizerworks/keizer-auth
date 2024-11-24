package repositories

import (
	"fmt"
	"keizer-auth/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	error := r.
		db.Model(&user).
		Clauses(clause.Returning{}).
		Create(user).
		Error
	return error
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

func (r *UserRepository) GetUserByStruct(user *models.User) error {
	result := r.db.Where(user).First(user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil
		}
		return fmt.Errorf("error in getting user: %w", result.Error)
	}

	return nil
}

func (r *UserRepository) UpdateUser(id string, user *models.User) error {
	return r.db.Model(user).Clauses(clause.Returning{}).Where("id = ?", id).Updates(user).Error
}
