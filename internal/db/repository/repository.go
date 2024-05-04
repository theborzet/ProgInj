package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/theborzet/library_project/internal/db/models"
)

type Repository interface {
	DeleteRecord(tableName string, id uint) error
	GetBookID(id uint) (*models.Book, error)
	GetAllBooks(genre, title string, authorID uint, yearFrom, yearTo int) ([]*models.Book, error)
	UpdateBook(id uint, updated *models.Book) error
	AddBook(book *models.Book) error
	GetAuthorID(id uint) (*models.Author, error)
	GetAllAuthors(first_name, last_name string) ([]*models.Author, error)
	UpdateAuthor(id uint, updated *models.Author) error
	AddAuthor(author *models.Author) error
	GetClientID(id uint) (*models.Client, error)
	AddClient(client *models.Client) error
	GetAuthorBooks(author_id uint) ([]*models.Book, error)
	GetAllGenres() ([]*string, error)
	GetUser(username string) (*models.Client, error)
	UserExists(username string) error
	ViewClientBook(client_id int) ([]*models.Book, error)
	AddClientBook(client_id, book_id int) error
}

type SQLRepository struct {
	db *sqlx.DB
}

func NewSQLRepository(db *sqlx.DB) *SQLRepository {
	return &SQLRepository{
		db: db,
	}
}

func (r *SQLRepository) DeleteRecord(tableName string, id uint) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableName)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *SQLRepository) GetBookID(id uint) (*models.Book, error) {
	query := "SELECT id, title, author_id, publication_year, genre, description, photo_url FROM book WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var book models.Book
	if err := row.Scan(&book.ID, &book.Title, &book.AuthorID, &book.PublicationYear, &book.Genre, &book.Description, &book.ImageUrl); err != nil {
		return nil, err
	}
	return &book, nil
}
func (r *SQLRepository) GetAllBooks(genre, title string, authorID uint, yearFrom, yearTo int) ([]*models.Book, error) {
	query := "SELECT id, title, author_id, publication_year, genre, description, photo_url FROM book WHERE 1=1 ORDER BY id ASC"
	if title != "" {
		lowercaseTitle := strings.ToLower(title)
		query += fmt.Sprintf(" AND LOWER(title) LIKE '%%%s%%'", lowercaseTitle) // Используем оператор LIKE для поиска по названию
	}
	if genre != "" {
		query += fmt.Sprintf(" AND genre = '%s'", genre)
	}
	if authorID != 0 {
		query += fmt.Sprintf(" AND author_id = %d", authorID)
	}
	if yearFrom != 0 {
		query += fmt.Sprintf(" AND publication_year >= %d", yearFrom)
	}
	if yearTo != 0 {
		query += fmt.Sprintf(" AND publication_year <= %d", yearTo)
	}
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	var books []*models.Book

	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.AuthorID, &book.PublicationYear, &book.Genre, &book.Description, &book.ImageUrl); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return books, nil
}

func (r *SQLRepository) UpdateBook(id uint, updated *models.Book) error {
	query := "UPDATE book SET title = $1, author_id = $2, publication_year = $3, genre = $4, description = $5, photo_url = $6 WHERE id = $7"
	_, err := r.db.Exec(query, updated.Title, updated.AuthorID, updated.PublicationYear, updated.Genre, updated.Description, updated.ImageUrl, id)
	return err
}

func (r *SQLRepository) AddBook(book *models.Book) error {
	query := "INSERT INTO book (title, author_id, publication_year, genre, description, photo_url) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := r.db.Exec(query, book.Title, book.AuthorID, book.PublicationYear, book.Genre, book.Description, book.ImageUrl)
	return err
}
func (r *SQLRepository) GetAuthorID(id uint) (*models.Author, error) {
	query := "SELECT id, first_name, last_name, birth_date, photo_url FROM author WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var author models.Author
	if err := row.Scan(&author.ID, &author.FirstName, &author.LastName, &author.BirthDate, &author.ImageUrl); err != nil {
		return nil, err
	}
	return &author, nil
}
func (r *SQLRepository) GetAllAuthors(first_name, last_name string) ([]*models.Author, error) {
	query := "SELECT id, first_name, last_name, birth_date, photo_url FROM author WHERE 1=1 ORDER BY id ASC"
	if first_name != "" {
		lowercaseFirstName := strings.ToLower(first_name)
		query += fmt.Sprintf(" AND LOWER(first_name) LIKE '%%%s%%' OR LOWER(last_name) LIKE '%%%s%%'", lowercaseFirstName, lowercaseFirstName) // Используем оператор LIKE для поиска по названию
	}
	if last_name != "" {
		lowercaseLastName := strings.ToLower(last_name)
		query += fmt.Sprintf(" AND LOWER(last_name) LIKE '%%%s%%' OR LOWER(first_name) LIKE '%%%s%%'", lowercaseLastName, lowercaseLastName) // Используем оператор LIKE для поиска по названию
	}
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	var authors []*models.Author

	for rows.Next() {
		var author models.Author
		if err := rows.Scan(&author.ID, &author.FirstName, &author.LastName, &author.BirthDate, &author.ImageUrl); err != nil {
			return nil, err
		}
		authors = append(authors, &author)
	}

	return authors, nil
}

func (r *SQLRepository) UpdateAuthor(id uint, updated *models.Author) error {
	query := "UPDATE author SET first_name = $1, last_name = $2, birth_date = $3, photo_url = $4 WHERE id = $5"
	_, err := r.db.Exec(query, updated.FirstName, updated.LastName, updated.BirthDate, updated.ImageUrl, id)
	return err
}

func (r *SQLRepository) AddAuthor(author *models.Author) error {
	query := "INSERT INTO author (first_name, last_name, birth_date, photo_url) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(query, author.FirstName, author.LastName, author.BirthDate, author.ImageUrl)
	return err
}

func (r *SQLRepository) GetAuthorBooks(author_id uint) ([]*models.Book, error) {
	query := "SELECT id, title, author_id, publication_year, genre, description, photo_url FROM book WHERE author_id=$1"
	rows, err := r.db.Query(query, author_id)

	if err != nil {
		return nil, err
	}

	var books []*models.Book

	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.AuthorID, &book.PublicationYear, &book.Genre, &book.Description, &book.ImageUrl); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return books, nil
}

func (r *SQLRepository) GetClientID(id uint) (*models.Client, error) {
	query := "SELECT id, username, password, email, access_level FROM client WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var client models.Client
	if err := row.Scan(&client.ID, &client.Username, &client.Password, &client.Email, &client.AccessLevel); err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *SQLRepository) AddClient(client *models.Client) error {
	query := "INSERT INTO client (username, password, email, access_level) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(query, client.Username, client.Password, client.Email, client.AccessLevel)
	return err // Возвращаем ошибку, если запрос выполнен с ошибкой
}

func (r *SQLRepository) ViewClientBook(client_id int) ([]*models.Book, error) {
	query := "SELECT DISTINCT b.id, b.title, b.author_id, b.publication_year, b.genre, b.description, b.photo_url FROM book AS b JOIN client_book AS c ON b.id = c.book_id WHERE c.client_id = $1 ORDER BY b.id ASC"
	rows, err := r.db.Query(query, client_id)
	if err != nil {
		return nil, err
	}

	var books []*models.Book

	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.AuthorID, &book.PublicationYear, &book.Genre, &book.Description, &book.ImageUrl); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, nil
}
func (r *SQLRepository) AddClientBook(client_id, book_id int) error {
	query := "INSERT INTO client_book (client_id, book_id) VALUES ($1, $2)"

	_, err := r.db.Exec(query, client_id, book_id)
	return err
}

func (r *SQLRepository) GetAllGenres() ([]*string, error) {
	query := "SELECT DISTINCT genre FROM book"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	var genres []*string

	for rows.Next() {
		var genre string
		if err := rows.Scan(&genre); err != nil {
			return nil, err
		}
		genres = append(genres, &genre)
	}

	return genres, nil
}

func (r *SQLRepository) GetUser(username string) (*models.Client, error) {
	query := "SELECT id, password, access_level FROM client WHERE username = $1"
	row := r.db.QueryRow(query, username)

	var NewClient models.Client

	if err := row.Scan(&NewClient.ID, &NewClient.Password, &NewClient.AccessLevel); err != nil {
		return nil, err
	}
	return &NewClient, nil

}

func (r *SQLRepository) UserExists(username string) error {

	query := "SELECT EXISTS(SELECT 1 FROM client WHERE username = ?)"

	var exists bool
	if err := r.db.QueryRow(query, username).Scan(&exists); err != nil {
		return err
	}

	return nil
}
