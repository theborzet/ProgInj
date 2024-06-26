package db

import (
	"fmt"
	"log"

	"github.com/theborzet/library_project/pkg/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Init(c *config.Config) *sqlx.DB {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.DBPort, c.User, c.Password, c.DBname)

	db, err := sqlx.Open("postgres", url)

	if err != nil {
		log.Fatalln(err)
	}

	migrateDB(db)

	return db
}

func migrateDB(db *sqlx.DB) {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS author (
            id SERIAL PRIMARY KEY,
            first_name VARCHAR(100) NOT NULL,
            last_name VARCHAR(100) NOT NULL,
            birth_date DATE,
			photo_url TEXT
        )
    `)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS book (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			author_id INTEGER REFERENCES Author(id),
			publication_year INT,
			genre VARCHAR(100),
			photo_url TEXT,
			description TEXT
		)
    `)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS client (
            id SERIAL PRIMARY KEY,
            username VARCHAR(100) NOT NULL UNIQUE,
            password VARCHAR(100) NOT NULL,
			email VARCHAR(100) NOT NULL,
            access_level INT DEFAULT 0
        )
    `)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS client_book (
		client_id INTEGER REFERENCES client(id),
		book_id INTEGER REFERENCES book(id),
		PRIMARY KEY (client_id, book_id)
		)
    `)
	if err != nil {
		log.Fatalln(err)
	}
}
