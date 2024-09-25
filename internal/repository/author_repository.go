package repository

import (
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	DB *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{
		DB: db,
	}
}

func (r *AuthorRepository) GetAll() ([]domain.Author, error) {
	var authors []domain.Author
	err := r.DB.Find(&authors).Error
	return authors, err
}

func (r *AuthorRepository) GetByID(id uint) (*domain.Author, error) {
	var author domain.Author
	err := r.DB.First(&author, id).Error
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *AuthorRepository) Create(author *domain.Author) error {
	return r.DB.Create(author).Error
}

func (r *AuthorRepository) Update(author *domain.Author) error {
	return r.DB.Save(author).Error
}

func (r *AuthorRepository) Delete(id uint) error {
	return r.DB.Delete(&domain.Author{}, id).Error
}
