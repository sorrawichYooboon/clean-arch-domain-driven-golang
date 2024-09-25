package domain

type Book struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear int    `json:"published_year"`
	Category      string `json:"category"`
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
