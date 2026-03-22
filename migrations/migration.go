package migrations

import (
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/infrastructure/database"
	"gorm.io/gorm"
)

func MigrateTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&database.Author{},
		&database.Book{},
		&database.User{},
	)
}
