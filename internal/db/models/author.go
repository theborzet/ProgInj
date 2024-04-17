package models

type Author struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	BirthDate string `json:"birth_date" db:"birth_date"`
}
