package repository

import (
	"backend/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindByID(id uuid.UUID) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *entity.User) error {
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	return r.db.Create(user).Error
}

func (r *userRepository) FindByID(id uuid.UUID) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
