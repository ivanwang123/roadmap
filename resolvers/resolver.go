package resolvers

import (
	"github.com/ivanwang123/roadmap/internal/checkpoint"
	"github.com/ivanwang123/roadmap/internal/checkpoint_status"
	"github.com/ivanwang123/roadmap/internal/roadmap"
	"github.com/ivanwang123/roadmap/internal/roadmap_follower"
	"github.com/ivanwang123/roadmap/internal/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CheckpointUsecase       checkpoint.Usecase
	CheckpointStatusUsecase checkpoint_status.Usecase
	RoadmapUsecase          roadmap.Usecase
	RoadmapFollowerUsecase  roadmap_follower.Usecase
	UserUsecase             user.Usecase
}
