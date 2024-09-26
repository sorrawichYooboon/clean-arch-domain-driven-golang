package domain

type Book struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear int    `json:"published_year"`
	Category      string `json:"category"`
}

func NewBook(title, author, category string, publishedYear int) *Book {
	return &Book{
		Title:         title,
		Author:        author,
		PublishedYear: publishedYear,
		Category:      category,
	}
}
