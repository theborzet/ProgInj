package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/db/models"
)

func (h Handler) AddBook(c *fiber.Ctx) error {
	title := c.FormValue("title")
	author := c.FormValue("author")
	author_id, err := strconv.ParseUint(author, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	publication_year := c.FormValue("publication_year")
	publication_year_id, err := strconv.ParseUint(publication_year, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	genre := c.FormValue("genre")
	description := c.FormValue("description")
	image_url := c.FormValue("image_url")

	book := models.Book{
		Title:           title,
		AuthorID:        uint(author_id),
		PublicationYear: uint(publication_year_id),
		Genre:           genre,
		Description:     description,
		ImageUrl:        image_url,
	}

	errchan := make(chan error)
	go func() {
		err := h.repo.AddBook(&book)
		if err != nil {
			errchan <- fiber.NewError(fiber.StatusBadRequest, err.Error())
		} else {
			errchan <- nil
		}
	}()

	err = <-errchan
	if err != nil {
		return err
	}

	return c.Redirect("/books", fiber.StatusFound)
}
