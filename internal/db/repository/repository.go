package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/theborzet/library_project/internal/db/models"
)

type Repository interface {
	DeleteRecord(tableName string, id int) error
	GetBookID(id int) (*models.Book, error)
	GetAllBooks() ([]*models.Book, error)
	UpdateBook(id int, updated *models.Book) error
	AddBook(book *models.Book) error
	GetAuthorID(id int) (*models.Author, error)
	GetAllAuthors() ([]*models.Author, error)
	UpdateAuthor(id int, updated *models.Author) error
	AddAuthor(author *models.Author) error
	GetClientID(id int) (*models.Client, error)
	GetAllClients() ([]*models.Client, error)
	UpdateClient(id int, updated *models.Client) error
	AddClient(client *models.Client) error
	// SignInUser(client *models.Client) error
	// LogInUser(client *models.Client) error
}

type SQLRepository struct {
	db *sqlx.DB
}

func NewSQLRepository(db *sqlx.DB) *SQLRepository {
	return &SQLRepository{
		db: db,
	}
}

func (r *SQLRepository) DeleteRecord(tableName string, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableName)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *SQLRepository) GetBookID(id int) (*models.Book, error) {
	query := "SELECT id, title, author_id, publication_year, genre FROM book WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var book models.Book
	if err := row.Scan(&book.ID, &book.Title, &book.AuthorID, &book.PublicationYear, &book.Genre); err != nil {
		return nil, err
	}
	return &book, nil
}
func (r *SQLRepository) GetAllBooks() ([]*models.Book, error) {
	query := "SELECT id, title, author_id, publication_year, genre FROM book"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	var books []*models.Book

	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.AuthorID, &book.PublicationYear, &book.Genre); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return books, nil
}

func (r *SQLRepository) UpdateBook(id int, updated *models.Book) error {
	query := "UPDATE book SET title = $1, author_id = $2, publication_year = $3, genre = $4 WHERE id = $5"
	_, err := r.db.Exec(query, updated.Title, updated.AuthorID, updated.PublicationYear, updated.Genre, id)
	return err
}

func (r *SQLRepository) AddBook(book *models.Book) error {
	query := "INSERT INTO book (title, author_id, publication_year, genre) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(query, book.Title, book.AuthorID, book.PublicationYear, book.Genre)
	return err
}
func (r *SQLRepository) GetAuthorID(id int) (*models.Author, error) {
	query := "SELECT id, first_name, last_name, birth_date FROM author WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var author models.Author
	if err := row.Scan(&author.ID, &author.FirstName, &author.LastName, &author.BirthDate); err != nil {
		return nil, err
	}
	return &author, nil
}
func (r *SQLRepository) GetAllAuthors() ([]*models.Author, error) {
	query := "SELECT id, first_name, last_name, birth_date FROM author"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	var authors []*models.Author

	for rows.Next() {
		var author models.Author
		if err := rows.Scan(&author.ID, &author.FirstName, &author.LastName, &author.BirthDate); err != nil {
			return nil, err
		}
		authors = append(authors, &author)
	}

	return authors, nil
}

func (r *SQLRepository) UpdateAuthor(id int, updated *models.Author) error {
	query := "UPDATE author SET first_name = $1, last_name = $2, birth_date = $3 WHERE id = $4"
	_, err := r.db.Exec(query, updated.FirstName, updated.LastName, updated.BirthDate, id)
	return err
}

func (r *SQLRepository) AddAuthor(author *models.Author) error {
	query := "INSERT INTO author (first_name, last_name, birth_date) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(query, author.FirstName, author.LastName, author.BirthDate)
	return err
}
func (r *SQLRepository) GetClientID(id int) (*models.Client, error) {
	query := "SELECT id, username, password,, email, access_level, books FROM client WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var client models.Client
	if err := row.Scan(&client.ID, &client.Username, &client.Password, &client.Email, &client.AccessLevel, &client.Books); err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *SQLRepository) GetAllClients() ([]*models.Client, error) {
	query := "SELECT id, username, password, email, access_level, books FROM client"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	var clients []*models.Client

	for rows.Next() {
		var client models.Client
		if err := rows.Scan(&client.ID, &client.Username, &client.Password, &client.Email, &client.AccessLevel, &client.Books); err != nil {
			return nil, err
		}
		clients = append(clients, &client)
	}

	return clients, nil
}

func (r *SQLRepository) UpdateClient(id int, updated *models.Client) error {
	query := "UPDATE client SET username = $1, password = $2, email = $3 access_level = $4, books = $5 WHERE id = $6"
	_, err := r.db.Exec(query, updated.Username, updated.Password, updated.Email, updated.AccessLevel, updated.Books, id)
	return err
}

func (r *SQLRepository) AddClient(client *models.Client) error {
	query := "INSERT INTO client (username, password, email, access_level, books) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.Exec(query, client.Username, client.Password, client.Email, client.AccessLevel, client.Books)
	return err
}
