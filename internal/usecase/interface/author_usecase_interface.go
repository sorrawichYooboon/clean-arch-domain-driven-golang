package usecaseinterface

import "github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"

type AuthorUseCase interface {
	GetAllAuthors() ([]domain.Author, error)
	GetAuthorByID(id uint) (*domain.Author, error)
	CreateAuthor(author *domain.Author) error
	UpdateAuthor(author *domain.Author) error
	DeleteAuthor(id uint) error
}
