package domain

type Author struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
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
