package usecase

import (
	"fmt"

	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
)

type BookUseCase struct {
	Repo      domain.BookRepository
	CacheRepo domain.CacheBookRepository
}

func NewBookUseCase(repo domain.BookRepository, cacheRepo domain.CacheBookRepository) *BookUseCase {
	return &BookUseCase{
		Repo:      repo,
		CacheRepo: cacheRepo,
	}
}

func (uc *BookUseCase) GetAllBooks() ([]domain.Book, error) {
	books, err := uc.CacheRepo.GetAll()
	if err == nil && books != nil {
		return books, nil
	}

	books, err = uc.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	if err := uc.CacheRepo.SetAll(books); err != nil {
		return nil, err
	}

	return books, nil
}

func (uc *BookUseCase) GetBookByID(id uint) (*domain.Book, error) {
	return uc.Repo.GetByID(id)
}

func (uc *BookUseCase) CreateBook(book *domain.Book) error {
	err := uc.Repo.Create(book)
	if err != nil {
		return err
	}

	books, cacheErr := uc.CacheRepo.GetAll()
	if cacheErr == nil {
		books = append(books, *book)
		if cacheErr := uc.CacheRepo.SetAll(books); cacheErr != nil {
			fmt.Println("Failed to update cache:", cacheErr)
		}
	}
	return nil
}

func (uc *BookUseCase) UpdateBook(book *domain.Book) error {
	err := uc.Repo.Update(book)
	if err != nil {
		return err
	}

	books, cacheErr := uc.CacheRepo.GetAll()
	if cacheErr == nil {
		for i, b := range books {
			if b.ID == book.ID {
				books[i] = *book
				break
			}
		}
		if cacheErr := uc.CacheRepo.SetAll(books); cacheErr != nil {
			fmt.Println("Failed to update cache:", cacheErr)
		}
	}
	return nil
}

func (uc *BookUseCase) DeleteBook(id uint) error {
	err := uc.Repo.Delete(id)
	if err != nil {
		return err
	}

	books, cacheErr := uc.CacheRepo.GetAll()
	if cacheErr == nil {
		for i, b := range books {
			if b.ID == id {
				books = append(books[:i], books[i+1:]...)
				break
			}
		}
		if cacheErr := uc.CacheRepo.SetAll(books); cacheErr != nil {
			fmt.Println("Failed to update cache:", cacheErr)
		}
	}
	return nil
}
