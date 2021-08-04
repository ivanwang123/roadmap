package roadmap

import (
	"context"

	"github.com/ivanwang123/roadmap/models"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*models.Roadmap, error)
	GetByID(ctx context.Context, ID int) (*models.Roadmap, error)
	GetByCreatorID(ctx context.Context, creatorID int) ([]*models.Roadmap, error)
	GetByPagination(ctx context.Context, input *models.GetRoadmaps) ([]*models.Roadmap, error)
	Create(ctx context.Context, input *models.NewRoadmap) (*models.Roadmap, error)
}

type Usecase interface {
	GetAll(ctx context.Context) ([]*models.Roadmap, error)
	GetByID(ctx context.Context, ID int) (*models.Roadmap, error)
	GetByCreatorID(ctx context.Context, creatorID int) ([]*models.Roadmap, error)
	GetByPagination(ctx context.Context, input *models.GetRoadmaps) ([]*models.Roadmap, error)
	Create(ctx context.Context, input *models.NewRoadmap) (*models.Roadmap, error)
}
