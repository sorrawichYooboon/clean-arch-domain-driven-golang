package database

import (
	"time"

	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/usecase"
	"gorm.io/gorm"
)

type Book struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title         string    `gorm:"type:varchar(255);not null" json:"title"`
	Author        string    `gorm:"type:varchar(255);not null" json:"author"`
	PublishedYear int       `gorm:"type:int;not null" json:"published_year"`
	Category      string    `gorm:"type:varchar(50)" json:"category"`
	CreatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func mapToDomainBook(b Book) domain.Book {
	return domain.Book{
		ID:            b.ID,
		Title:         b.Title,
		Author:        b.Author,
		PublishedYear: b.PublishedYear,
		Category:      b.Category,
	}
}

func mapToDBBook(b domain.Book) Book {
	return Book{
		ID:            b.ID,
		Title:         b.Title,
		Author:        b.Author,
		PublishedYear: b.PublishedYear,
		Category:      b.Category,
	}
}

type BookRepositoryImpl struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) usecase.BookRepository {
	return &BookRepositoryImpl{
		DB: db,
	}
}

func (r *BookRepositoryImpl) GetAll() ([]domain.Book, error) {
	var dbBooks []Book
	err := r.DB.Find(&dbBooks).Error
	if err != nil {
		return nil, err
	}
	var books []domain.Book
	for _, b := range dbBooks {
		books = append(books, mapToDomainBook(b))
	}
	return books, nil
}

func (r *BookRepositoryImpl) GetByID(id uint) (*domain.Book, error) {
	var book Book
	err := r.DB.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	domainBook := mapToDomainBook(book)
	return &domainBook, nil
}

func (r *BookRepositoryImpl) Create(book *domain.Book) error {
	dbBook := mapToDBBook(*book)
	err := r.DB.Create(&dbBook).Error
	if err == nil {
		book.ID = dbBook.ID
	}
	return err
}

func (r *BookRepositoryImpl) Update(book *domain.Book) error {
	dbBook := mapToDBBook(*book)
	return r.DB.Save(&dbBook).Error
}

func (r *BookRepositoryImpl) Delete(id uint) error {
	return r.DB.Delete(&Book{}, id).Error
}
