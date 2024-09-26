package dto

type AuthorCreateDTO struct {
	Name string `json:"name" validate:"required"`
	Bio  string `json:"bio" validate:"required"`
}
