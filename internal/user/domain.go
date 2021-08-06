package user

import (
	"context"

	"github.com/ivanwang123/roadmap/internal/common/transaction"
	"github.com/ivanwang123/roadmap/models"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*models.User, error)
	GetByID(ctx context.Context, ID int) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	GetByRoadmapFollowing(ctx context.Context, roadmapID int) ([]*models.User, error)
	GetIn(ctx context.Context, IDs []string) ([]*models.User, error)
	Create(ctx context.Context, input *models.NewUser) (*models.User, error)
	WithTransaction(ctx context.Context, fn transaction.TxFunc) error
}

type Usecase interface {
	GetAll(ctx context.Context) ([]*models.User, error)
	GetByID(ctx context.Context, ID int) (*models.User, error)
	GetByRoadmapFollowing(ctx context.Context, roadmapID int) ([]*models.User, error)
	Create(ctx context.Context, input *models.NewUser) (*models.User, error)
	Authenticate(ctx context.Context, input *models.Login) (*models.User, error)
	BatchGet(ctx context.Context) func(int) (*models.User, error)
}
