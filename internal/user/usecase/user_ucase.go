package usecase

import (
	"context"

	"github.com/ivanwang123/roadmap/internal/user"
	"github.com/ivanwang123/roadmap/models"
)

type userUsecase struct {
	userRepo user.Repository
}

func NewUserUsecase(userRepo user.Repository) user.Usecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) GetAll(ctx context.Context) ([]*models.User, error) {
	return u.userRepo.GetAll(ctx)
}

func (u *userUsecase) GetByID(ctx context.Context, ID int) (*models.User, error) {
	return u.userRepo.GetByID(ctx, ID)
}

func (u *userUsecase) Create(ctx context.Context, input *models.NewUser) (*models.User, error) {
	return u.userRepo.Create(ctx, input)
}

func (u *userUsecase) Authenticate(ctx context.Context, input *models.Login) (*models.User, error) {
	return u.userRepo.Authenticate(ctx, input)
}
