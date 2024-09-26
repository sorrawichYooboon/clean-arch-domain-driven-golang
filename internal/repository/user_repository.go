package repository

import "github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindByUsername(username string) (*domain.User, error)
}
