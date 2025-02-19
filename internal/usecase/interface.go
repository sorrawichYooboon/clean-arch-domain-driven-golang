package usecase

import "github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"

type AuthorUseCase interface {
	GetAllAuthors() ([]domain.Author, error)
	GetAuthorByID(id uint) (*domain.Author, error)
	CreateAuthor(name, bio string) error
	UpdateAuthor(author *domain.Author) error
	DeleteAuthor(id uint) error
}

type BookUseCase interface {
	GetAllBooks() ([]domain.Book, error)
	GetBookByID(id uint) (*domain.Book, error)
	CreateBook(title, author, category string, publishedYear int) error
	UpdateBook(book *domain.Book) error
	DeleteBook(id uint) error
}

type UserUseCase interface {
	CreateUser(username, password string) error
	Authenticate(username, password string) (*domain.User, error)
}
