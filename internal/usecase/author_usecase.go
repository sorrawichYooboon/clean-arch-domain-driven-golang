package usecase

import (
	"fmt"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/pkg/apperror"
)

type AuthorUseCaseImpl struct {
	authorRepo AuthorRepository
	cacheRepo  CacheAuthorRepository
}

func NewAuthorUseCase(authorRepo AuthorRepository, cacheRepo CacheAuthorRepository) AuthorUseCase {
	return &AuthorUseCaseImpl{
		authorRepo: authorRepo,
		cacheRepo:  cacheRepo,
	}
}

func (uc *AuthorUseCaseImpl) GetAllAuthors() ([]domain.Author, error) {
	authors, err := uc.cacheRepo.GetAll()
	if err != nil {
		return nil, &apperror.ErrCacheDatabase
	}

	if len(authors) > 0 {
		return authors, nil
	}

	authors, err = uc.authorRepo.GetAll()
	if err != nil {
		return nil, &apperror.ErrDatabase
	}

	if err := uc.cacheRepo.SetAll(authors); err != nil {
		return nil, &apperror.ErrCacheDatabase
	}

	return authors, nil
}

func (uc *AuthorUseCaseImpl) GetAuthorByID(id uint) (*domain.Author, error) {
	author, err := uc.authorRepo.GetByID(id)
	if err != nil {
		return nil, &apperror.ErrDatabase
	}
	return author, nil
}

func (uc *AuthorUseCaseImpl) CreateAuthor(name, bio string) error {
	author := domain.NewAuthor(name, bio)
	err := uc.authorRepo.Create(author)
	if err != nil {
		return &apperror.ErrDatabase
	}

	authors, cacheErr := uc.cacheRepo.GetAll()
	if cacheErr == nil {
		authors = append(authors, *author)
		if cacheErr := uc.cacheRepo.SetAll(authors); cacheErr != nil {
			fmt.Println("Failed to update cache:", cacheErr)
		}
	}
	return nil
}

func (uc *AuthorUseCaseImpl) UpdateAuthor(author *domain.Author) error {
	err := uc.authorRepo.Update(author)
	if err != nil {
		return &apperror.ErrDatabase
	}

	authors, cacheErr := uc.cacheRepo.GetAll()
	if cacheErr == nil {
		for i, a := range authors {
			if a.ID == author.ID {
				authors[i] = *author
				break
			}
		}
		if cacheErr := uc.cacheRepo.SetAll(authors); cacheErr != nil {
			fmt.Println("Failed to update cache:", cacheErr)
		}
	}
	return nil
}

func (uc *AuthorUseCaseImpl) DeleteAuthor(id uint) error {
	err := uc.authorRepo.Delete(id)
	if err != nil {
		return &apperror.ErrDatabase
	}

	authors, cacheErr := uc.cacheRepo.GetAll()
	if cacheErr == nil {
		for i, a := range authors {
			if a.ID == id {
				authors = append(authors[:i], authors[i+1:]...)
				break
			}
		}
		if cacheErr := uc.cacheRepo.SetAll(authors); cacheErr != nil {
			fmt.Println("Failed to update cache:", cacheErr)
		}
	}
	return nil
}
