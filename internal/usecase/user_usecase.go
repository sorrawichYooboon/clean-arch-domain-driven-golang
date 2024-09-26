package usecase

import (
	"errors"

	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

func (uc *UserUseCase) CreateUser(username, password string) error {
	existingUser, _ := uc.userRepo.FindByUsername(username)
	if existingUser != nil {
		return errors.New("username already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := domain.NewUser(username, string(hashedPassword))
	return uc.userRepo.Create(user)
}

func (uc *UserUseCase) Authenticate(username, password string) (*domain.User, error) {
	user, err := uc.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}
