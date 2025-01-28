package repository

import "time"

type Book struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title         string    `gorm:"type:varchar(255);not null" json:"title"`
	Author        string    `gorm:"type:varchar(255);not null" json:"author"`
	PublishedYear int       `gorm:"type:int;not null" json:"published_year"`
	Category      string    `gorm:"type:varchar(50)" json:"category"`
	CreatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type BookRepository interface {
	GetAll() ([]Book, error)
	GetByID(id uint) (*Book, error)
	Create(book *Book) error
	Update(book *Book) error
	Delete(id uint) error
}

type CacheBookRepository interface {
	GetAll() ([]Book, error)
	SetAll(books []Book) error
}
