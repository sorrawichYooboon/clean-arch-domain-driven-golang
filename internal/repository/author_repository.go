package repository

import "github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"

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
