package database

import (
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

func (r *BookRepositoryImpl) GetAll() ([]repository.Book, error) {
	var books []repository.Book
	err := r.DB.Find(&books).Error
	return books, err
}

func (r *BookRepositoryImpl) GetByID(id uint) (*repository.Book, error) {
	var book repository.Book
	err := r.DB.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepositoryImpl) Create(book *repository.Book) error {
	return r.DB.Create(book).Error
}

func (r *BookRepositoryImpl) Update(book *repository.Book) error {
	return r.DB.Save(book).Error
}

func (r *BookRepositoryImpl) Delete(id uint) error {
	return r.DB.Delete(&repository.Book{}, id).Error
}
