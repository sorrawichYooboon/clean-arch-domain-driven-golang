package database

import (
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/repository"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(user *repository.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepositoryImpl) FindByUsername(username string) (*repository.User, error) {
	var user repository.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
