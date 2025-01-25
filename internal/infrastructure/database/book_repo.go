package database

import (
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/repository"
	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) repository.BookRepository {
	return &BookRepositoryImpl{
		DB: db,
	}
}

func (r *BookRepositoryImpl) GetAll() ([]domain.Book, error) {
	var books []domain.Book
	err := r.DB.Find(&books).Error
	return books, err
}

func (r *BookRepositoryImpl) GetByID(id uint) (*domain.Book, error) {
	var book domain.Book
	err := r.DB.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepositoryImpl) Create(book *domain.Book) error {
	return r.DB.Create(book).Error
}

func (r *BookRepositoryImpl) Update(book *domain.Book) error {
	return r.DB.Save(book).Error
}

func (r *BookRepositoryImpl) Delete(id uint) error {
	return r.DB.Delete(&domain.Book{}, id).Error
}
