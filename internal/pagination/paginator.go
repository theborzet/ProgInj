package pagination

import (
	"fmt"

	"github.com/theborzet/library_project/internal/db/models"
)

func PaginateBooks(books []*models.Book, page, pageSize int) ([]*models.Book, Paginator) {
	var paginator Paginator

	totalItems := len(books)
	if totalItems == 0 {
		return nil, Paginator{}
	}
	paginator.TotalItems = totalItems
	paginator.PageSize = pageSize
	paginator.TotalPages = (totalItems + pageSize - 1) / pageSize

	if page < 1 {
		page = 1
	} else if page > paginator.TotalPages {
		page = paginator.TotalPages
	}

	paginator.CurrentPage = page

	start := (page - 1) * pageSize
	end := start + pageSize
	if end > totalItems {
		end = totalItems
	}

	paginatedBooks := books[start:end]

	paginator.HasPrevious = page > 1
	paginator.HasNext = page < paginator.TotalPages

	if paginator.HasPrevious {
		paginator.PreviousPage = page - 1
	}

	if paginator.HasNext {
		paginator.NextPage = page + 1
	}

	paginator.PageNumbers = make([]PageNumber, paginator.TotalPages)
	for i := range paginator.PageNumbers {
		paginator.PageNumbers[i] = PageNumber{
			Number: i + 1,
			URL:    fmt.Sprintf("/books?page=%d", i+1),
		}
	}

	return paginatedBooks, paginator
}

func PaginateAuthors(authors []*models.Author, page, pageSize int) ([]*models.Author, Paginator) {
	var paginator Paginator

	totalItems := len(authors)
	if totalItems == 0 {
		return nil, Paginator{}
	}
	paginator.TotalItems = totalItems
	paginator.PageSize = pageSize
	paginator.TotalPages = (totalItems + pageSize - 1) / pageSize

	if page < 1 {
		page = 1
	} else if page > paginator.TotalPages {
		page = paginator.TotalPages
	}

	paginator.CurrentPage = page

	start := (page - 1) * pageSize
	end := start + pageSize
	if end > totalItems {
		end = totalItems
	}

	paginatedAuthors := authors[start:end]

	paginator.HasPrevious = page > 1
	paginator.HasNext = page < paginator.TotalPages

	if paginator.HasPrevious {
		paginator.PreviousPage = page - 1
	}

	if paginator.HasNext {
		paginator.NextPage = page + 1
	}

	paginator.PageNumbers = make([]PageNumber, paginator.TotalPages)
	for i := range paginator.PageNumbers {
		paginator.PageNumbers[i] = PageNumber{
			Number: i + 1,
			URL:    fmt.Sprintf("/author?page=%d", i+1),
		}
	}

	return paginatedAuthors, paginator
}
