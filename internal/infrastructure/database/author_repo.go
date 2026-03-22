package database

import (
	"time"

	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/usecase"
	"gorm.io/gorm"
)

type Author struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Bio       string    `gorm:"type:text" json:"bio"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func mapToDomainAuthor(a Author) domain.Author {
	return domain.Author{
		ID:   a.ID,
		Name: a.Name,
		Bio:  a.Bio,
	}
}

func mapToDBAuthor(a domain.Author) Author {
	return Author{
		ID:   a.ID,
		Name: a.Name,
		Bio:  a.Bio,
	}
}

type AuthorRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) usecase.AuthorRepository {
	return &AuthorRepositoryImpl{
		DB: db,
	}
}

func (r *AuthorRepositoryImpl) GetAll() ([]domain.Author, error) {
	var dbAuthors []Author
	err := r.DB.Find(&dbAuthors).Error
	if err != nil {
		return nil, err
	}
	var authors []domain.Author
	for _, a := range dbAuthors {
		authors = append(authors, mapToDomainAuthor(a))
	}
	return authors, err
}

func (r *AuthorRepositoryImpl) GetByID(id uint) (*domain.Author, error) {
	var author Author
	err := r.DB.First(&author, id).Error
	if err != nil {
		return nil, err
	}
	domainAuthor := mapToDomainAuthor(author)
	return &domainAuthor, nil
}

func (r *AuthorRepositoryImpl) Create(author *domain.Author) error {
	dbAuthor := mapToDBAuthor(*author)
	err := r.DB.Create(&dbAuthor).Error
	if err == nil {
		author.ID = dbAuthor.ID
	}
	return err
}

func (r *AuthorRepositoryImpl) Update(author *domain.Author) error {
	dbAuthor := mapToDBAuthor(*author)
	return r.DB.Save(&dbAuthor).Error
}

func (r *AuthorRepositoryImpl) Delete(id uint) error {
	return r.DB.Delete(&Author{}, id).Error
}
