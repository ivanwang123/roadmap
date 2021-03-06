package postgres

import (
	"context"

	"github.com/ivanwang123/roadmap/internal/common/transaction"
	"github.com/ivanwang123/roadmap/internal/roadmap"
	"github.com/ivanwang123/roadmap/internal/roadmap/repository"
	"github.com/ivanwang123/roadmap/models"
	"github.com/jmoiron/sqlx"
)

type roadmapRepo struct {
	db *sqlx.DB
}

func NewRoadmapRepo(db *sqlx.DB) roadmap.Repository {
	return &roadmapRepo{db}
}

func (r *roadmapRepo) GetAll(ctx context.Context) ([]*models.Roadmap, error) {
	roadmaps := []*models.Roadmap{}
	conn := transaction.GetConn(ctx, r.db)
	if err := conn.Select(&roadmaps, "SELECT * FROM roadmaps"); err != nil {
		return nil, err
	}
	return roadmaps, nil
}

func (r *roadmapRepo) GetByID(ctx context.Context, ID int) (*models.Roadmap, error) {
	var roadmap models.Roadmap
	conn := transaction.GetConn(ctx, r.db)
	if err := conn.Get(&roadmap, "SELECT * FROM roadmaps WHERE id = $1 LIMIT 1", ID); err != nil {
		return nil, err
	}
	return &roadmap, nil
}

func (r *roadmapRepo) GetByCreatorID(ctx context.Context, creatorID int) ([]*models.Roadmap, error) {
	roadmaps := []*models.Roadmap{}
	conn := transaction.GetConn(ctx, r.db)
	if err := conn.Select(&roadmaps, "SELECT * FROM roadmaps WHERE creator_id = $1", creatorID); err != nil {
		return nil, err
	}
	return roadmaps, nil
}

func (r *roadmapRepo) GetByFollower(ctx context.Context, userID int) ([]*models.Roadmap, error) {
	roadmaps := []*models.Roadmap{}
	conn := transaction.GetConn(ctx, r.db)
	if err := conn.Select(&roadmaps, "SELECT r.* FROM roadmap_followers AS f LEFT JOIN roadmaps AS r ON f.roadmap_id = r.id WHERE f.user_id = $1", userID); err != nil {
		return nil, err
	}
	return roadmaps, nil
}

func (r *roadmapRepo) GetByPagination(ctx context.Context, input *models.GetRoadmaps) ([]*models.Roadmap, error) {
	roadmaps := []*models.Roadmap{}

	var query string
	switch input.Sort {
	case models.SortNewest:
		query = "SELECT * FROM roadmaps WHERE (created_at, id) <= ($1, $2) ORDER BY created_at DESC, id DESC LIMIT $3"
	case models.SortOldest:
		query = "SELECT * FROM roadmaps WHERE (created_at, id) >= ($1, $2) ORDER BY created_at ASC, id ASC LIMIT $3"
	case models.SortMostFollowers:
		query = "SELECT r.* FROM roadmaps r LEFT JOIN roadmap_followers rf ON rf.roadmap_id = r.id GROUP BY r.id HAVING (count(rf.roadmap_id), r.id) <= ($1, $2) ORDER BY count(rf.roadmap_id) DESC, r.id DESC LIMIT $3"
	case models.SortMostCheckpoints:
		query = "SELECT r.* FROM roadmaps r LEFT JOIN checkpoints c ON c.roadmap_id = r.id GROUP BY r.id HAVING (count(c.id), r.id) <= ($1, $2) ORDER BY count(c.id) DESC, r.id DESC LIMIT $3"
	case models.SortLeastCheckpoints:
		query = "SELECT r.* FROM roadmaps r LEFT JOIN checkpoints c ON c.roadmap_id = r.id GROUP BY r.id HAVING (count(c.id), r.id) >= ($1, $2) ORDER BY count(c.id) ASC, r.id ASC LIMIT $3"
	}

	conn := transaction.GetConn(ctx, r.db)
	if err := conn.Select(&roadmaps, query, input.CursorValue, input.CursorID, repository.PaginationLimit); err != nil {
		return nil, err
	}
	return roadmaps, nil
}

func (r *roadmapRepo) GetIn(ctx context.Context, IDs []string) ([]*models.Roadmap, error) {
	roadmaps := []*models.Roadmap{}
	query, args, err := sqlx.In("SELECT * FROM roadmaps WHERE id IN (?)", IDs)
	if err != nil {
		return nil, err
	}

	conn := transaction.GetConn(ctx, r.db)
	if err := conn.Select(&roadmaps, r.db.Rebind(query), args...); err != nil {
		return nil, err
	}

	return roadmaps, nil
}

func (r *roadmapRepo) Create(ctx context.Context, input *models.NewRoadmap) (*models.Roadmap, error) {
	var roadmap models.Roadmap
	conn := transaction.GetConn(ctx, r.db)
	if err := conn.Get(&roadmap, "INSERT INTO roadmaps (title, description, creator_id) VALUES ($1, $2, $3) RETURNING *",
		input.Title, input.Description, input.CreatorID); err != nil {
		return nil, err
	}
	return &roadmap, nil
}

func (r *roadmapRepo) WithTransaction(ctx context.Context, fn transaction.TxFunc) error {
	return transaction.NewTransaction(ctx, r.db, fn)
}
