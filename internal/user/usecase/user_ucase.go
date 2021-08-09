package usecase

import (
	"context"
	"errors"
	"strconv"

	"github.com/graph-gophers/dataloader"
	"github.com/ivanwang123/roadmap/internal/common/loader"
	"github.com/ivanwang123/roadmap/internal/user"
	"github.com/ivanwang123/roadmap/internal/user/repository"
	"github.com/ivanwang123/roadmap/models"
)

type userUsecase struct {
	userRepo user.Repository
}

func NewUserUsecase(u user.Repository) user.Usecase {
	return &userUsecase{
		userRepo: u,
	}
}

func (u *userUsecase) GetAll(ctx context.Context) ([]*models.User, error) {
	return u.userRepo.GetAll(ctx)
}

func (u *userUsecase) GetByID(ctx context.Context, ID int) (*models.User, error) {
	return u.userRepo.GetByID(ctx, ID)
}

func (u *userUsecase) GetByRoadmapFollowing(ctx context.Context, roadmapID int) ([]*models.User, error) {
	return u.userRepo.GetByRoadmapFollowing(ctx, roadmapID)
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
		err = errors.New("missing credentials")
	}

	if err != nil {
		return nil, err
	}

	if !repository.CheckPasswordHash(user.Password, input.Password) {
		return nil, errors.New("incorrect password")
	}

	return user, nil
}

func (u *userUsecase) BatchGet(ctx context.Context) func(int) (*models.User, error) {
	batchedLoader := loader.NewLoader(func(ctx context.Context, keys []string) (map[string]interface{}, error) {
		users, err := u.userRepo.GetIn(ctx, keys)
		if err != nil {
			return nil, err
		}

		usersMap := make(map[string]interface{})
		for _, user := range users {
			usersMap[strconv.Itoa(user.ID)] = user
		}
		return usersMap, nil
	})

	return func(userID int) (*models.User, error) {
		result, err := batchedLoader.Load(ctx, dataloader.StringKey(strconv.Itoa(userID)))()
		if err != nil {
			return nil, err
		}

		user, _ := result.(*models.User)
		return user, nil
	}
}
