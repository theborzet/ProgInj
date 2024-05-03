package handlers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/db/models"
)

func (h Handler) UpdateAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	updatedId, err := strconv.Atoi(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	name := c.FormValue("name")
	surname := c.FormValue("surname")
	birth_date := c.FormValue("birth_date")
	layout := "2006-01-02"
	birthDate, err := time.Parse(layout, birth_date)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	image_url := c.FormValue("photo_url")

	author := models.Author{
		FirstName: name,
		LastName:  surname,
		BirthDate: birthDate,
		ImageUrl:  image_url,
	}

	errchan := make(chan error)

	go func() {
		if err := h.repo.UpdateAuthor(uint(updatedId), &author); err != nil {
			errchan <- err
		} else {
			errchan <- nil
		}

	}()
	err = <-errchan
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Redirect("/authors", fiber.StatusFound)
}
