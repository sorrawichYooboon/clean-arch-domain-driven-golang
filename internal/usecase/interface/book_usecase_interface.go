package usecaseinterface

import "github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"

type BookUseCase interface {
	GetAllBooks() ([]domain.Book, error)
	GetBookByID(id uint) (*domain.Book, error)
	CreateBook(title, author, category string, publishedYear int) error
	UpdateBook(book *domain.Book) error
	DeleteBook(id uint) error
}
