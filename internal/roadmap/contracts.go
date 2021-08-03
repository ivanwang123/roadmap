package roadmap

import (
	"context"

	"github.com/ivanwang123/roadmap/internal/roadmap/repository"
	"github.com/ivanwang123/roadmap/models"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*models.Roadmap, error)
	GetByID(ctx context.Context, ID int) (*models.Roadmap, error)
	GetByCreatorID(ctx context.Context, creatorID int) ([]*models.Roadmap, error)
	GetByPagination(ctx context.Context, input repository.PaginationInput) ([]*models.Roadmap, error)
}

type Usecase interface {
	GetAll(ctx context.Context) ([]*models.Roadmap, error)
	GetByID(ctx context.Context, ID int) (*models.Roadmap, error)
	GetByCreatorID(ctx context.Context, creatorID int) ([]*models.Roadmap, error)
	GetByPagination(ctx context.Context, input repository.PaginationInput) ([]*models.Roadmap, error)
}
