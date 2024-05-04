package handlers

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/pagination"
)

func (h Handler) ViewAllAuthors(c *fiber.Ctx) error {
	sess, err := h.session.Get(c)
	if err != nil {
		return err
	}
	isAuthenticated, ok := sess.Get("isAuthenticated").(bool)
	if !ok {
		isAuthenticated = false
	}
	userID, ok := sess.Get("uID").(int)
	if !ok {
		userID = 0
	}
	access_level, ok := sess.Get("acessLevel").(int)
	if !ok {
		userID = 0
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}
	fullname := c.FormValue("fullname")
	var first_name, last_name string
	parts := strings.Split(fullname, " ")
	if len(parts) == 1 {
		first_name = parts[0]
	} else if len(parts) >= 2 {
		first_name = parts[0]
		last_name = parts[1]
	}
	authors, err := h.repo.GetAllAuthors(first_name, last_name)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	pageSize := 4
	paginatedAuthors, paginator := pagination.PaginateAuthors(authors, page, pageSize)

	return c.Render("author_list", fiber.Map{
		"ClientId":        userID,
		"Paginator":       paginator,
		"Title":           "Список авторов",
		"IsPaginated":     paginator.TotalPages > 1,
		"Authors":         paginatedAuthors,
		"IsAuthenticated": isAuthenticated,
		"Access_level":    access_level,
	})
}
