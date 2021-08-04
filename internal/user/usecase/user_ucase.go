package usecase

import (
	"context"
	"errors"

	"github.com/ivanwang123/roadmap/internal/user"
	"github.com/ivanwang123/roadmap/internal/user/repository"
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
	hashedPassword, err := repository.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	input.Password = hashedPassword

	return u.userRepo.Create(ctx, input)
}

func (u *userUsecase) Authenticate(ctx context.Context, input *models.Login) (*models.User, error) {
	var user *models.User
	var err error
	if input.Email != nil {
		user, err = u.userRepo.GetByEmail(ctx, *input.Email)
	} else if input.Username != nil {
		user, err = u.userRepo.GetByUsername(ctx, *input.Username)
	} else {
		err = errors.New("Missing credentials")
	}

	if err != nil {
		return nil, err
	}

	if !repository.CheckPasswordHash(user.Password, input.Password) {
		return nil, errors.New("Incorrect password")
	}

	return user, nil
}
