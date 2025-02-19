package dto

type GetAllBookResponse struct {
	ID            uint   `json:"id"`
	Title         string `json:"title" validate:"required"`
	Author        string `json:"author" validate:"required"`
	PublishedYear int    `json:"published_year" validate:"required"`
	Category      string `json:"category" validate:"required"`
}

type CreateBookRequest struct {
	Title         string `json:"title" validate:"required"`
	Author        string `json:"author" validate:"required"`
	PublishedYear int    `json:"published_year" validate:"required"`
	Category      string `json:"category" validate:"required"`
}

type UpdateBookRequest struct {
	Title         string `json:"title" validate:"required"`
	Author        string `json:"author" validate:"required"`
	PublishedYear int    `json:"published_year" validate:"required"`
	Category      string `json:"category" validate:"required"`
}
