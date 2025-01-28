package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/repository"
	usecaseinterface "github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/usecase/interface"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCaseImpl struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) usecaseinterface.UserUseCase {
	return &UserUseCaseImpl{userRepo: userRepo}
}

func (uc *UserUseCaseImpl) CreateUser(username, password string) error {
	existingUser, _ := uc.userRepo.FindByUsername(username)
	if existingUser != nil {
		return errors.New("username already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := domain.NewUser(username, string(hashedPassword))

	createUserReq := mapToRepoUser(*user)

	return uc.userRepo.Create(&createUserReq)
}

func (uc *UserUseCaseImpl) Authenticate(username, password string) (*domain.User, error) {
	user, err := uc.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	resp := mapToDomainUser(*user)

	return &resp, nil
}

func mapToDomainUser(repoUser repository.User) domain.User {
	id, _ := uuid.Parse(repoUser.ID)
	return domain.User{
		ID:       id,
		Username: repoUser.Username,
		Password: repoUser.Password,
	}
}

func mapToRepoUser(domainUser domain.User) repository.User {
	return repository.User{
		ID:       domainUser.ID.String(),
		Username: domainUser.Username,
		Password: domainUser.Password,
	}
}
