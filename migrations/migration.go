package migrations

import (
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/repository"
	"gorm.io/gorm"
)

func MigrateTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&repository.Author{},
		&repository.Book{},
		&repository.User{},
	)
}
