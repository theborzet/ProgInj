package models

type Client struct {
	ID               int    `json:"id" db:"id"`
	Username         string `json:"username" db:"username"`
	Password         string `json:"password" db:"password"`
	Email            string `json:"email" db:"email"`
	AccessLevel      int    `json:"access_level" db:"access_level"`
	Books            []Book `json:"books" db:"books"`
	IsAuthentificate bool   `json:"is_authentificate"`
}

func NewClient() *Client {
	return &Client{
		IsAuthentificate: false,
	}
}
