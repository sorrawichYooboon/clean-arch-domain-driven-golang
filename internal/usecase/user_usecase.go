package usecase

import (
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/pkg/apperror"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCaseImpl struct {
	userRepo UserRepository
}

func NewUserUseCase(userRepo UserRepository) UserUseCase {
	return &UserUseCaseImpl{userRepo: userRepo}
}

func (uc *UserUseCaseImpl) CreateUser(username, password string) error {
	existingUser, _ := uc.userRepo.FindByUsername(username)
	if existingUser != nil {
		return &apperror.ErrUserExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return &apperror.ErrHashPassword
	}

	user := domain.NewUser(username, string(hashedPassword))
	err = uc.userRepo.Create(user)
	if err != nil {
		return &apperror.ErrDatabase
	}

	return nil
}

func (uc *UserUseCaseImpl) Authenticate(username, password string) (*domain.User, error) {
	user, err := uc.userRepo.FindByUsername(username)
	if err != nil {
		return nil, &apperror.ErrDatabase
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, &apperror.ErrInvalidCredentials
	}

	return user, nil
}
