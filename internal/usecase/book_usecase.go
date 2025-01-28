package usecase

import (
	"fmt"

	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/repository"
	usecaseinterface "github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/usecase/interface"
)

type BookUseCaseImpl struct {
	bookRepo  repository.BookRepository
	cacheRepo repository.CacheBookRepository
}

func NewBookUseCase(bookRepo repository.BookRepository, cacheRepo repository.CacheBookRepository) usecaseinterface.BookUseCase {
	return &BookUseCaseImpl{
		bookRepo:  bookRepo,
		cacheRepo: cacheRepo,
	}
}

func (uc *BookUseCaseImpl) GetAllBooks() ([]domain.Book, error) {
	resp := []domain.Book{}
	books, err := uc.cacheRepo.GetAll()
	if err != nil {
		return nil, err
	}

	if len(books) > 0 {
		for _, book := range books {
			resp = append(resp, mapToDomainBook(book))
		}
	}

	books, err = uc.bookRepo.GetAll()
	if err != nil {
		return nil, err
	}

	if err := uc.cacheRepo.SetAll(books); err != nil {
		return nil, err
	}

	for _, book := range books {
		resp = append(resp, mapToDomainBook(book))
	}

	return resp, nil
}

func (uc *BookUseCaseImpl) GetBookByID(id uint) (*domain.Book, error) {
	book, err := uc.bookRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	resp := mapToDomainBook(*book)

	return &resp, nil
}

func (uc *BookUseCaseImpl) CreateBook(title, author, category string, publishedYear int) error {
	book := domain.NewBook(title, author, category, publishedYear)

	createBookReq := &repository.Book{
		Title:         book.Title,
		Author:        book.Author,
		Category:      book.Category,
		PublishedYear: book.PublishedYear,
	}

	err := uc.bookRepo.Create(createBookReq)
	if err != nil {
		return err
	}

	books, cacheErr := uc.cacheRepo.GetAll()
	if cacheErr == nil {
		cacheBook := mapToRepoBook(*book)

		books = append(books, cacheBook)
		if cacheErr := uc.cacheRepo.SetAll(books); cacheErr != nil {
			fmt.Println("Failed to update cache:", cacheErr)
		}
	}
	return nil
}

func (uc *BookUseCaseImpl) UpdateBook(book *domain.Book) error {
	updateBookReq := &repository.Book{
		ID:            book.ID,
		Title:         book.Title,
		Category:      book.Category,
		PublishedYear: book.PublishedYear,
	}

	err := uc.bookRepo.Update(updateBookReq)
	if err != nil {
		return err
	}

	books, cacheErr := uc.cacheRepo.GetAll()
	if cacheErr == nil {
		for i, b := range books {
			if b.ID == book.ID {
				cacheBook := mapToRepoBook(*book)

				books[i] = cacheBook
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
		return err
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

func mapToDomainBook(repoBook repository.Book) domain.Book {
	return domain.Book{
		ID:            repoBook.ID,
		Title:         repoBook.Title,
		Author:        repoBook.Author,
		PublishedYear: repoBook.PublishedYear,
		Category:      repoBook.Category,
	}
}

func mapToRepoBook(domainBook domain.Book) repository.Book {
	return repository.Book{
		ID:            domainBook.ID,
		Title:         domainBook.Title,
		Author:        domainBook.Author,
		PublishedYear: domainBook.PublishedYear,
		Category:      domainBook.Category,
	}
}
