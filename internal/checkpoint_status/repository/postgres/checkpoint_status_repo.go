package postgres

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/ivanwang123/roadmap/internal/checkpoint_status"
	"github.com/ivanwang123/roadmap/models"
	"github.com/jmoiron/sqlx"
)

type checkpointStatusRepo struct {
	db *sqlx.DB
}

func NewCheckpointStatusRepo(db *sqlx.DB) checkpoint_status.Repository {
	return &checkpointStatusRepo{db}
}

func (r *checkpointStatusRepo) Get(ctx context.Context, input *models.GetCheckpointStatus) (*models.CheckpointStatus, error) {
	var checkpointStatus models.CheckpointStatus
	if err := r.db.Get(&checkpointStatus, "SELECT * FROM checkpoint_status WHERE user_id = $1 AND checkpoint_id = $2 AND roadmap_id = $3 LIMIT 1", input.UserID, input.CheckpointID, input.RoadmapID); err != nil {
		return nil, err
	}
	return &checkpointStatus, nil
}

func (r *checkpointStatusRepo) Update(ctx context.Context, userID int, input *models.UpdateStatus) error {
	if _, err := r.db.Exec("UPDATE checkpoint_status SET status = $1 WHERE user_id = $2 AND checkpoint_id = $3", input.Status, userID, input.CheckpointID); err != nil {
		return err
	}
	return nil
}

func (r *checkpointStatusRepo) CreateMany(ctx context.Context, input []*models.CreateCheckpointStatus) error {
	if len(input) == 0 {
		return nil
	}

	values := []string{}
	for _, status := range input {
		values = append(values, fmt.Sprintf("(%d, %d, %d)", status.UserID, status.CheckpointID, status.RoadmapID))
	}
	query := fmt.Sprintf("INSERT INTO checkpoint_status (user_id, checkpoint_id, roadmap_id) VALUES %s", strings.Join(values, ", "))

	if _, err := r.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func (r *checkpointStatusRepo) DeleteMany(ctx context.Context, input *models.DeleteManyCheckpointStatus) error {
	var userQuery string
	if len(input.UserIDs) == 1 {
		userQuery = fmt.Sprintf("= %d", input.UserIDs[0])
	} else {
		strUserIds := make([]string, len(input.UserIDs))
		for i, id := range input.UserIDs {
			strUserIds[i] = strconv.Itoa(id)
		}
		userQuery = fmt.Sprintf("IN (%s)", strings.Join(strUserIds, ", "))
	}

	var checkpointQuery string
	if len(input.CheckpointIDs) == 1 {
		checkpointQuery = fmt.Sprintf("= %d", input.CheckpointIDs[0])
	} else {
		strCheckpointIds := make([]string, len(input.CheckpointIDs))
		for i, id := range input.CheckpointIDs {
			strCheckpointIds[i] = strconv.Itoa(id)
		}
		checkpointQuery = fmt.Sprintf("IN (%s)", strings.Join(strCheckpointIds, ", "))
	}

	query := fmt.Sprintf("DELETE FROM checkpoint_status WHERE roadmap_id = %d AND user_id %s AND checkpoint_id %s", input.RoadmapID, userQuery, checkpointQuery)
	if _, err := r.db.Exec(query); err != nil {
		return err
	}
	return nil
}
