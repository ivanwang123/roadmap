package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ivanwang123/roadmap/graphql/generated"
	"github.com/ivanwang123/roadmap/internal/common/auth"
	"github.com/ivanwang123/roadmap/models"
)

func (r *queryResolver) User(ctx context.Context, input *models.GetUser) (*models.User, error) {
	return r.UserUsecase.GetByID(ctx, input.ID)
}

func (r *queryResolver) Roadmap(ctx context.Context, input *models.GetRoadmap) (*models.Roadmap, error) {
	return r.RoadmapUsecase.GetByID(ctx, input.ID)
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return r.UserUsecase.GetAll(ctx)
}

func (r *queryResolver) Roadmaps(ctx context.Context, input *models.GetRoadmaps) ([]*models.Roadmap, error) {
	return r.RoadmapUsecase.GetByPagination(ctx, input)
}

func (r *queryResolver) Me(ctx context.Context) (*models.User, error) {
	userID, err := auth.GetCurrentUser(ctx)
	if err != nil {
		return nil, nil
	}
	return r.UserUsecase.GetByID(ctx, userID)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
