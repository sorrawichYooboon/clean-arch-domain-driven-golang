package usecase

import (
	"fmt"

	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/infrastructure/cache"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/repository"
)

type AuthorUseCase struct {
	Repo      repository.AuthorRepository
	CacheRepo cache.CacheAuthorRepository
}

func NewAuthorUseCase(repo repository.AuthorRepository, cacheRepo cache.CacheAuthorRepository) *AuthorUseCase {
	return &AuthorUseCase{
		Repo:      repo,
		CacheRepo: cacheRepo,
	}
}

func (uc *AuthorUseCase) GetAllAuthors() ([]domain.Author, error) {
	authors, err := uc.CacheRepo.GetAll()
	if err == nil && authors != nil {
		return authors, nil
	}

	authors, err = uc.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	if err := uc.CacheRepo.SetAll(authors); err != nil {
		fmt.Println("Failed to update cache:", err)
	}

	return authors, nil
}

func (uc *AuthorUseCase) GetAuthorByID(id uint) (*domain.Author, error) {
	return uc.Repo.GetByID(id)
}

func (uc *AuthorUseCase) CreateAuthor(name, bio string) error {
	author := domain.NewAuthor(name, bio)
	err := uc.Repo.Create(author)
	if err != nil {
		return err
	}

	authors, cacheErr := uc.CacheRepo.GetAll()
	if cacheErr == nil {
		authors = append(authors, *author)
		if cacheErr := uc.CacheRepo.SetAll(authors); cacheErr != nil {
			fmt.Println("Failed to update cache:", cacheErr)
		}
	}
	return nil
}

func (uc *AuthorUseCase) UpdateAuthor(author *domain.Author) error {
	err := uc.Repo.Update(author)
	if err != nil {
		return err
	}

	authors, cacheErr := uc.CacheRepo.GetAll()
	if cacheErr == nil {
		for i, a := range authors {
			if a.ID == author.ID {
				authors[i] = *author
				break
			}
		}
		if cacheErr := uc.CacheRepo.SetAll(authors); cacheErr != nil {
			fmt.Println("Failed to update cache:", cacheErr)
		}
	}
	return nil
}

func (uc *AuthorUseCase) DeleteAuthor(id uint) error {
	err := uc.Repo.Delete(id)
	if err != nil {
		return err
	}

	authors, cacheErr := uc.CacheRepo.GetAll()
	if cacheErr == nil {
		for i, a := range authors {
			if a.ID == id {
				authors = append(authors[:i], authors[i+1:]...)
				break
			}
		}
		if cacheErr := uc.CacheRepo.SetAll(authors); cacheErr != nil {
			fmt.Println("Failed to update cache:", cacheErr)
		}
	}
	return nil
}
