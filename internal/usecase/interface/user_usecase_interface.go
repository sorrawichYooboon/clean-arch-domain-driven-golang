package usecaseinterface

import "github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"

type UserUseCase interface {
	CreateUser(username, password string) error
	Authenticate(username, password string) (*domain.User, error)
}
