package resolvers

import (
	"github.com/ivanwang123/roadmap/internal/checkpoint"
	checkpoint_postgres "github.com/ivanwang123/roadmap/internal/checkpoint/repository/postgres"
	checkpoint_usecase "github.com/ivanwang123/roadmap/internal/checkpoint/usecase"
	"github.com/ivanwang123/roadmap/internal/checkpoint_status"
	checkpoint_status_postgres "github.com/ivanwang123/roadmap/internal/checkpoint_status/repository/postgres"
	checkpoint_status_usecase "github.com/ivanwang123/roadmap/internal/checkpoint_status/usecase"
	"github.com/ivanwang123/roadmap/internal/roadmap"
	roadmap_postgres "github.com/ivanwang123/roadmap/internal/roadmap/repository/postgres"
	roadmap_usecase "github.com/ivanwang123/roadmap/internal/roadmap/usecase"
	"github.com/ivanwang123/roadmap/internal/roadmap_follower"
	roadmap_follower_postgres "github.com/ivanwang123/roadmap/internal/roadmap_follower/repository/postgres"
	roadmap_follower_usecase "github.com/ivanwang123/roadmap/internal/roadmap_follower/usecase"
	"github.com/ivanwang123/roadmap/internal/user"
	user_postgres "github.com/ivanwang123/roadmap/internal/user/repository/postgres"
	user_usecase "github.com/ivanwang123/roadmap/internal/user/usecase"
	"github.com/jmoiron/sqlx"
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

func NewResolver(db *sqlx.DB) *Resolver {
	checkpointRepo := checkpoint_postgres.NewCheckpointRepo(db)
	checkpointStatusRepo := checkpoint_status_postgres.NewCheckpointStatusRepo(db)
	roadmapRepo := roadmap_postgres.NewRoadmapRepo(db)
	roadmapFollowerRepo := roadmap_follower_postgres.NewRoadmapFollowerRepo(db)
	userRepo := user_postgres.NewUserRepo(db)

	checkpointUsecase := checkpoint_usecase.NewCheckpointUsecase(checkpointRepo, checkpointStatusRepo, roadmapFollowerRepo)
	checkpointStatusUsecase := checkpoint_status_usecase.NewCheckpointStatusUsecase(checkpointStatusRepo, checkpointRepo)
	roadmapUsecase := roadmap_usecase.NewRoadmapUsecase(roadmapRepo, roadmapFollowerRepo, checkpointRepo, checkpointStatusRepo)
	roadmapFollowerUsecase := roadmap_follower_usecase.NewRoadmapFollowerUsecase(roadmapFollowerRepo)
	userUsecase := user_usecase.NewUserUsecase(userRepo)

	return &Resolver{
		CheckpointUsecase:       checkpointUsecase,
		CheckpointStatusUsecase: checkpointStatusUsecase,
		RoadmapUsecase:          roadmapUsecase,
		RoadmapFollowerUsecase:  roadmapFollowerUsecase,
		UserUsecase:             userUsecase,
	}
}
