package repository

import (
	"time"
)

type Author struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Bio       string    `gorm:"type:text" json:"bio"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type AuthorRepository interface {
	GetAll() ([]Author, error)
	GetByID(id uint) (*Author, error)
	Create(author *Author) error
	Update(author *Author) error
	Delete(id uint) error
}

type CacheAuthorRepository interface {
	GetAll() ([]Author, error)
	SetAll(authors []Author) error
}
