package repository

import "github.com/ivanwang123/roadmap/models"

const PaginationLimit = 10

type PaginationInput struct {
	CursorID    int
	CursorValue string
	SortBy      models.Sort
}
