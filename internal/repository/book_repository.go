package repository

import "github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"

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
