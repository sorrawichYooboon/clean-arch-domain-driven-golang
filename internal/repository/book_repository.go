package repository

import (
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		DB: db,
	}
}

func (r *BookRepository) GetAll() ([]domain.Book, error) {
	var books []domain.Book
	err := r.DB.Find(&books).Error
	return books, err
}

func (r *BookRepository) GetByID(id uint) (*domain.Book, error) {
	var book domain.Book
	err := r.DB.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) Create(book *domain.Book) error {
	return r.DB.Create(book).Error
}

func (r *BookRepository) Update(book *domain.Book) error {
	return r.DB.Save(book).Error
}

func (r *BookRepository) Delete(id uint) error {
	return r.DB.Delete(&domain.Book{}, id).Error
}
