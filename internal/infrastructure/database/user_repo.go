package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/domain"
	"github.com/sorrawichYooboon/clean-arch-domain-driven-golang/internal/usecase"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Username  string    `gorm:"type:varchar(255);not null;unique" json:"username"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func mapToDomainUser(u User) domain.User {
	id, _ := uuid.Parse(u.ID)
	return domain.User{
		ID:       id,
		Username: u.Username,
		Password: u.Password,
	}
}

func mapToDBUser(u domain.User) User {
	return User{
		ID:       u.ID.String(),
		Username: u.Username,
		Password: u.Password,
	}
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) usecase.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(user *domain.User) error {
	dbUser := mapToDBUser(*user)
	err := r.db.Create(&dbUser).Error
	if err == nil {
		id, _ := uuid.Parse(dbUser.ID)
		user.ID = id
	}
	return err
}

func (r *UserRepositoryImpl) FindByUsername(username string) (*domain.User, error) {
	var user User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	domainUser := mapToDomainUser(user)
	return &domainUser, nil
}
