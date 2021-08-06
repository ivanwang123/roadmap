package roadmap

import (
	"context"

	"github.com/ivanwang123/roadmap/internal/common/transaction"
	"github.com/ivanwang123/roadmap/models"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*models.Roadmap, error)
	GetByID(ctx context.Context, ID int) (*models.Roadmap, error)
	GetByCreatorID(ctx context.Context, creatorID int) ([]*models.Roadmap, error)
	GetByFollower(ctx context.Context, userID int) ([]*models.Roadmap, error)
	GetByPagination(ctx context.Context, input *models.GetRoadmaps) ([]*models.Roadmap, error)
	GetIn(ctx context.Context, IDs []string) ([]*models.Roadmap, error)
	Create(ctx context.Context, input *models.NewRoadmap) (*models.Roadmap, error)
	WithTransaction(ctx context.Context, fn transaction.TxFunc) error
}

type Usecase interface {
	GetAll(ctx context.Context) ([]*models.Roadmap, error)
	GetByID(ctx context.Context, ID int) (*models.Roadmap, error)
	GetByCreatorID(ctx context.Context, creatorID int) ([]*models.Roadmap, error)
	GetByFollower(ctx context.Context, userID int) ([]*models.Roadmap, error)
	GetByPagination(ctx context.Context, input *models.GetRoadmaps) ([]*models.Roadmap, error)
	Create(ctx context.Context, input *models.NewRoadmap) (*models.Roadmap, error)
	ToggleFollow(ctx context.Context, userID, roadmapID int) (*models.Roadmap, error)
	BatchGet(ctx context.Context) func(int) (*models.Roadmap, error)
}
