package models

import "time"

type Author struct {
	ID        int       `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	BirthDate time.Time `json:"birth_date" db:"birth_date"`
	ImageUrl  string    `json:"image_url" db:"image_url"`
}
