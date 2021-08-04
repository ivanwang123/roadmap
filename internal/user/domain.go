package user

import (
	"context"

	"github.com/ivanwang123/roadmap/models"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*models.User, error)
	GetByID(ctx context.Context, ID int) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	Create(ctx context.Context, input *models.NewUser) (*models.User, error)
}

type Usecase interface {
	GetAll(ctx context.Context) ([]*models.User, error)
	GetByID(ctx context.Context, ID int) (*models.User, error)
	Create(ctx context.Context, input *models.NewUser) (*models.User, error)
	Authenticate(ctx context.Context, input *models.Login) (*models.User, error)
}
