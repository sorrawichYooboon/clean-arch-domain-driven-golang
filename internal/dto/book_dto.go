package dto

type BookDTO struct {
	Title         string `json:"title" validate:"required"`
	Author        string `json:"author" validate:"required"`
	PublishedYear int    `json:"published_year" validate:"required"`
	Category      string `json:"category" validate:"required"`
}
