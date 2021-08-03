package stores

import (
	"github.com/ivanwang123/roadmap/models"
	"github.com/jmoiron/sqlx"
)

type CheckpointStore struct {
	DB *sqlx.DB
}

// TODO: Put in transaction?
func (s *CheckpointStore) Create(input *models.NewCheckpoint) (*models.Checkpoint, error) {
	var checkpoint models.Checkpoint
	if err := s.DB.Get(&checkpoint, "INSERT INTO checkpoints (title, instructions, links, roadmap_id) VALUES ($1, $2, $3, $4) RETURNING *",
		input.Title, input.Instructions, input.Links, input.RoadmapID); err != nil {
		return nil, err
	}

	roadmapFollowers := []*models.RoadmapFollower{}
	if err := s.DB.Select(&roadmapFollowers, "SELECT * FROM roadmap_followers WHERE roadmap_id = $1", input.RoadmapID); err != nil {
		return nil, err
	}

	newCheckpointStatuses := make([]*NewCheckpointStatus, len(roadmapFollowers))
	for i, follower := range roadmapFollowers {
		newCheckpointStatuses[i] = &NewCheckpointStatus{userId: follower.UserID, checkpointId: checkpoint.ID, roadmapId: input.RoadmapID}
	}

	checkpointStatusStore := &CheckpointStatusStore{DB: s.DB}
	if err := checkpointStatusStore.CreateManyCheckpointStatus(newCheckpointStatuses); err != nil {
		return nil, err
	}

	return &checkpoint, nil
}

func (s *CheckpointStore) GetByRoadmap(roadmapId int) ([]*models.Checkpoint, error) {
	checkpoints := []*models.Checkpoint{}
	if err := s.DB.Select(&checkpoints, "SELECT * FROM checkpoints WHERE roadmap_id = $1",
		roadmapId); err != nil {
		return nil, err
	}
	return checkpoints, nil
}

func (s *CheckpointStore) GetCheckpointStatus(userId, checkpointId, roadmapId int) (*models.CheckpointStatus, error) {
	var checkpointStatus models.CheckpointStatus
	if err := s.DB.Get(&checkpointStatus, "SELECT * FROM checkpoint_status WHERE user_id = $1 AND checkpoint_id = $2 AND roadmap_id = $3 LIMIT 1", userId, checkpointId, roadmapId); err != nil {
		return nil, err
	}
	return &checkpointStatus, nil
}

func (s *CheckpointStore) UpdateStatus(userId int, input *models.UpdateStatus) (*models.Checkpoint, error) {
	if _, err := s.DB.Exec("UPDATE checkpoint_status SET status = $1 WHERE user_id = $2 AND checkpoint_id = $3", input.Status, userId, input.CheckpointID); err != nil {
		return nil, nil
	}

	var checkpoint models.Checkpoint
	if err := s.DB.Get(&checkpoint, "SELECT * FROM checkpoints WHERE id = $1", input.CheckpointID); err != nil {
		return nil, nil
	}
	checkpoint.Status = models.StatusType(input.Status)
	return &checkpoint, nil
}
