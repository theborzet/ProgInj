package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/db/models"
)

func (h Handler) AddAuthor(c *fiber.Ctx) error {
	name := c.FormValue("name")
	surname := c.FormValue("surname")
	birth_date := c.FormValue("birth_date")
	layout := "2006-01-02"
	birthDate, err := time.Parse(layout, birth_date)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	image_url := c.FormValue("image_url")

	author := models.Author{
		FirstName: name,
		LastName:  surname,
		BirthDate: birthDate,
		ImageUrl:  image_url,
	}

	errchan := make(chan error)
	go func() {
		err := h.repo.AddAuthor(&author)
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
	return c.Redirect("/authors", fiber.StatusFound)
}
