package database

import (
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/repository"
	"gorm.io/gorm"
)

type AuthorRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) repository.AuthorRepository {
	return &AuthorRepositoryImpl{
		DB: db,
	}
}

func (r *AuthorRepositoryImpl) GetAll() ([]repository.Author, error) {
	var authors []repository.Author
	err := r.DB.Find(&authors).Error
	return authors, err
}

func (r *AuthorRepositoryImpl) GetByID(id uint) (*repository.Author, error) {
	var author repository.Author
	err := r.DB.First(&author, id).Error
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *AuthorRepositoryImpl) Create(author *repository.Author) error {
	return r.DB.Create(author).Error
}

func (r *AuthorRepositoryImpl) Update(author *repository.Author) error {
	return r.DB.Save(author).Error
}

func (r *AuthorRepositoryImpl) Delete(id uint) error {
	return r.DB.Delete(&repository.Author{}, id).Error
}
