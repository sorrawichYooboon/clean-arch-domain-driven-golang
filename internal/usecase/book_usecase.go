package usecase

import (
	"fmt"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/pkg/apperror"
)

type BookUseCaseImpl struct {
	bookRepo  BookRepository
	cacheRepo CacheBookRepository
}

func NewBookUseCase(bookRepo BookRepository, cacheRepo CacheBookRepository) BookUseCase {
	return &BookUseCaseImpl{
		bookRepo:  bookRepo,
		cacheRepo: cacheRepo,
	}
}

func (uc *BookUseCaseImpl) GetAllBooks() ([]domain.Book, error) {
	books, err := uc.cacheRepo.GetAll()
	if err != nil {
		return nil, &apperror.ErrCacheDatabase
	}

	if len(books) > 0 {
		return books, nil
	}

	books, err = uc.bookRepo.GetAll()
	if err != nil {
		return nil, &apperror.ErrDatabase
	}

	if err := uc.cacheRepo.SetAll(books); err != nil {
		return nil, &apperror.ErrCacheDatabase
	}

	return books, nil
}

func (uc *BookUseCaseImpl) GetBookByID(id uint) (*domain.Book, error) {
	book, err := uc.bookRepo.GetByID(id)
	if err != nil {
		return nil, &apperror.ErrDatabase
	}
	return book, nil
}

func (uc *BookUseCaseImpl) CreateBook(title, author, category string, publishedYear int) error {
	book := domain.NewBook(title, author, category, publishedYear)
	err := uc.bookRepo.Create(book)
	if err != nil {
		return &apperror.ErrDatabase
	}

	books, cacheErr := uc.cacheRepo.GetAll()
	if cacheErr == nil {
		books = append(books, *book)
		if cacheErr := uc.cacheRepo.SetAll(books); cacheErr != nil {
			fmt.Println("Failed to update cache:", cacheErr)
		}
	}
	return nil
}

func (uc *BookUseCaseImpl) UpdateBook(book *domain.Book) error {
	err := uc.bookRepo.Update(book)
	if err != nil {
		return &apperror.ErrDatabase
	}

	books, cacheErr := uc.cacheRepo.GetAll()
	if cacheErr == nil {
		for i, b := range books {
			if b.ID == book.ID {
				books[i] = *book
				break
			}
		}
		if cacheErr := uc.cacheRepo.SetAll(books); cacheErr != nil {
			fmt.Println("Failed to update cache:", cacheErr)
		}
	}
	return nil
}

func (uc *BookUseCaseImpl) DeleteBook(id uint) error {
	err := uc.bookRepo.Delete(id)
	if err != nil {
		return &apperror.ErrDatabase
	}

	books, cacheErr := uc.cacheRepo.GetAll()
	if cacheErr == nil {
		for i, b := range books {
			if b.ID == id {
				books = append(books[:i], books[i+1:]...)
				break
			}
		}
		if cacheErr := uc.cacheRepo.SetAll(books); cacheErr != nil {
			fmt.Println("Failed to update cache:", cacheErr)
		}
	}
	return nil
}
