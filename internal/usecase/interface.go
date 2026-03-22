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

type BookRepository interface {
	GetAll() ([]domain.Book, error)
	GetByID(id uint) (*domain.Book, error)
	Create(book *domain.Book) error
	Update(book *domain.Book) error
	Delete(id uint) error
}

type CacheBookRepository interface {
	GetAll() ([]domain.Book, error)
	SetAll(books []domain.Book) error
}

type AuthorRepository interface {
	GetAll() ([]domain.Author, error)
	GetByID(id uint) (*domain.Author, error)
	Create(author *domain.Author) error
	Update(author *domain.Author) error
	Delete(id uint) error
}

type CacheAuthorRepository interface {
	GetAll() ([]domain.Author, error)
	SetAll(authors []domain.Author) error
}

type UserRepository interface {
	Create(user *domain.User) error
	FindByUsername(username string) (*domain.User, error)
}
