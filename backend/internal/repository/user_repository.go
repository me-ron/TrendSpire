package repository

import (
    "gorm.io/gorm"
    "backend/internal/entity"
)

type UserRepository interface {
    Create(user *entity.User) error
    FindByID(id uint) (*entity.User, error)
    FindByEmail(email string) (*entity.User, error)
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db}
}

func (r *userRepository) Create(user *entity.User) error {
    return r.db.Create(user).Error
}

func (r *userRepository) FindByID(id uint) (*entity.User, error) {
    var user entity.User
    if err := r.db.First(&user, id).Error; err != nil {
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
