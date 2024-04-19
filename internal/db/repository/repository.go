package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/theborzet/library_project/internal/db/models"
)

type Repository interface {
	GetProductID(id int) (*models.Book, error)
	//тут наши функции для хендлеров
}

type SQLRepository struct {
	db *sqlx.DB
}

func NewSQLRepository(db *sqlx.DB) *SQLRepository {
	return &SQLRepository{
		db: db,
	}
}

func (r *SQLRepository) GetProductID(id int) (*models.Book, error) {
	query := "SELECT id, title, author_id, publication_year, genre FROM book WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var book models.Book
	if err := row.Scan(&book.ID, &book.Title, &book.AuthorID, &book.PublicationYear, &book.Genre); err != nil {
		return nil, err
	}
	return &book, nil
}
