package postgres

import (
	"context"

	"github.com/ivanwang123/roadmap/internal/common/transaction"
	"github.com/ivanwang123/roadmap/internal/roadmap_follower"
	"github.com/ivanwang123/roadmap/models"
	"github.com/jmoiron/sqlx"
)

type roadmapFollowerRepo struct {
	db *sqlx.DB
}

func NewRoadmapFollowerRepo(db *sqlx.DB) roadmap_follower.Repository {
	return &roadmapFollowerRepo{db}
}

func (r *roadmapFollowerRepo) Get(ctx context.Context, userID int, roadmapID int) (*models.RoadmapFollower, error) {
	var roadmapFollower models.RoadmapFollower
	conn := transaction.GetConn(ctx, r.db)
	if err := conn.Get(&roadmapFollower, "SELECT * FROM roadmap_followers WHERE user_id = $1 AND roadmap_id = $2 LIMIT 1", userID, roadmapID); err != nil {
		return nil, err
	}
	return &roadmapFollower, nil
}

func (r *roadmapFollowerRepo) GetByRoadmap(ctx context.Context, roadmapID int) ([]*models.RoadmapFollower, error) {
	roadmapFollowers := []*models.RoadmapFollower{}
	conn := transaction.GetConn(ctx, r.db)
	if err := conn.Select(&roadmapFollowers, "SELECT * FROM roadmap_followers WHERE roadmap_id = $1", roadmapID); err != nil {
		return nil, err
	}
	return roadmapFollowers, nil
}

func (r *roadmapFollowerRepo) Create(ctx context.Context, userID, roadmapID int) error {
	conn := transaction.GetConn(ctx, r.db)
	if _, err := conn.Exec("INSERT INTO roadmap_followers (user_id, roadmap_id) VALUES ($1, $2)", userID, roadmapID); err != nil {
		return err
	}
	return nil
}

func (r *roadmapFollowerRepo) Delete(ctx context.Context, userID, roadmapID int) error {
	conn := transaction.GetConn(ctx, r.db)
	if _, err := conn.Exec("DELETE FROM roadmap_followers WHERE user_id = $1 AND roadmap_id = $2", userID, roadmapID); err != nil {
		return err
	}
	return nil
}

func (r *roadmapFollowerRepo) WithTransaction(ctx context.Context, fn transaction.TxFunc) error {
	return transaction.NewTransaction(ctx, r.db, fn)
}
