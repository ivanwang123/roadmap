package postgres

import (
	"context"

	"github.com/ivanwang123/roadmap/internal/checkpoint"
	"github.com/ivanwang123/roadmap/internal/common/transaction"
	"github.com/ivanwang123/roadmap/models"
	"github.com/jmoiron/sqlx"
)

type checkpointRepo struct {
	db *sqlx.DB
}

func NewCheckpointRepo(db *sqlx.DB) checkpoint.Repository {
	return &checkpointRepo{db}
}

func (r *checkpointRepo) GetByID(ctx context.Context, ID int) (*models.Checkpoint, error) {
	var checkpoint models.Checkpoint
	conn := transaction.GetConn(ctx, r.db)
	if err := conn.Get(&checkpoint, "SELECT * FROM checkpoints WHERE id = $1", ID); err != nil {
		return nil, err
	}
	return &checkpoint, nil
}

func (r *checkpointRepo) GetByRoadmap(ctx context.Context, roadmapID int) ([]*models.Checkpoint, error) {
	checkpoints := []*models.Checkpoint{}
	conn := transaction.GetConn(ctx, r.db)
	if err := conn.Select(&checkpoints, "SELECT * FROM checkpoints WHERE roadmap_id = $1",
		roadmapID); err != nil {
		return nil, err
	}
	return checkpoints, nil
}

func (r *checkpointRepo) GetIDByRoadmap(ctx context.Context, roadmapID int) ([]int, error) {
	checkpointIDs := []int{}
	conn := transaction.GetConn(ctx, r.db)
	err := conn.Select(&checkpointIDs, "SELECT id FROM checkpoints WHERE roadmap_id = $1", roadmapID)
	if err != nil {
		return nil, err
	}
	return checkpointIDs, nil
}

func (r *checkpointRepo) Create(ctx context.Context, input *models.NewCheckpoint) (*models.Checkpoint, error) {
	var checkpoint models.Checkpoint
	conn := transaction.GetConn(ctx, r.db)
	if err := conn.Get(&checkpoint, "INSERT INTO checkpoints (title, instructions, links, roadmap_id) VALUES ($1, $2, $3, $4) RETURNING *",
		input.Title, input.Instructions, input.Links, input.RoadmapID); err != nil {
		return nil, err
	}
	return &checkpoint, nil
}

func (r *checkpointRepo) WithTransaction(ctx context.Context, fn transaction.TxFunc) error {
	return transaction.NewTransaction(ctx, r.db, fn)
}
