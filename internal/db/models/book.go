package models

type Book struct {
	ID              uint   `json:"id" db:"id" `
	Title           string `json:"title" db:"title"`
	AuthorID        uint   `json:"author_id" db:"author_id" `
	PublicationYear uint   `json:"publication_year" db:"publication_year" `
	Genre           string `json:"genre" db:"genre"`
	Description     string `json:"description" db:"description"`
	ImageUrl        string `json:"image_url" db:"image_url"`
}
