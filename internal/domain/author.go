package domain

type Author struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

func NewAuthor(name, bio string) *Author {
	return &Author{
		Name: name,
		Bio:  bio,
	}
}
