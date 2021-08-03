package usecase

import (
	"context"

	"github.com/ivanwang123/roadmap/internal/roadmap"
	"github.com/ivanwang123/roadmap/internal/roadmap/repository"
	"github.com/ivanwang123/roadmap/models"
)

type roadmapUsecase struct {
	roadmapRepo roadmap.Repository
}

func NewRoadmapUsecase(roadmapRepo roadmap.Repository) roadmap.Usecase {
	return &roadmapUsecase{
		roadmapRepo: roadmapRepo,
	}
}

func (u *roadmapUsecase) GetAll(ctx context.Context) ([]*models.Roadmap, error) {
	return u.roadmapRepo.GetAll(ctx)
}

func (u *roadmapUsecase) GetByID(ctx context.Context, ID int) (*models.Roadmap, error) {
	return u.roadmapRepo.GetByID(ctx, ID)
}

func (u *roadmapUsecase) GetByCreatorID(ctx context.Context, creatorID int) ([]*models.Roadmap, error) {
	return u.roadmapRepo.GetByCreatorID(ctx, creatorID)
}

func (u *roadmapUsecase) GetByPagination(ctx context.Context, input repository.PaginationInput) ([]*models.Roadmap, error) {
	return u.roadmapRepo.GetByPagination(ctx, input)
}
