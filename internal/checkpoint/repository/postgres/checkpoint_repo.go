package postgres

import (
	"context"

	"github.com/ivanwang123/roadmap/internal/checkpoint"
	"github.com/ivanwang123/roadmap/models"
	"github.com/jmoiron/sqlx"
)

type checkpointRepo struct {
	db *sqlx.DB
}

func NewCheckpointRepo(db *sqlx.DB) checkpoint.Repository {
	return &checkpointRepo{db}
}

func (r *checkpointRepo) GetByRoadmap(ctx context.Context, roadmapID int) ([]*models.Checkpoint, error) {
	checkpoints := []*models.Checkpoint{}
	if err := r.db.Select(&checkpoints, "SELECT * FROM checkpoints WHERE roadmap_id = $1",
		roadmapID); err != nil {
		return nil, err
	}
	return checkpoints, nil
}

func (r *checkpointRepo) Create(ctx context.Context, input *models.NewCheckpoint) (*models.Checkpoint, error) {
	// TODO: Put in transaction?
	var checkpoint models.Checkpoint
	if err := r.db.Get(&checkpoint, "INSERT INTO checkpoints (title, instructions, links, roadmap_id) VALUES ($1, $2, $3, $4) RETURNING *",
		input.Title, input.Instructions, input.Links, input.RoadmapID); err != nil {
		return nil, err
	}

	roadmapFollowers := []*models.RoadmapFollower{}
	if err := r.db.Select(&roadmapFollowers, "SELECT * FROM roadmap_followers WHERE roadmap_id = $1", input.RoadmapID); err != nil {
		return nil, err
	}

	newCheckpointStatuses := make([]*models.NewCheckpointStatus, len(roadmapFollowers))
	for i, follower := range roadmapFollowers {
		newCheckpointStatuses[i] = &models.NewCheckpointStatus{UserID: follower.UserID, CheckpointID: checkpoint.ID, RoadmapID: input.RoadmapID}
	}

	checkpointStatusStore := &CheckpointStatusStore{DB: s.DB}
	if err := checkpointStatusStore.CreateManyCheckpointStatus(newCheckpointStatuses); err != nil {
		return nil, err
	}

	return &checkpoint, nil
}
