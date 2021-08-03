package stores

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

type CheckpointStatusStore struct {
	DB *sqlx.DB
}

type NewCheckpointStatus struct {
	userId       int
	roadmapId    int
	checkpointId int
}

func (s *CheckpointStatusStore) CreateManyCheckpointStatus(input []*NewCheckpointStatus) error {
	if len(input) == 0 {
		return nil
	}

	values := []string{}
	for _, status := range input {
		values = append(values, fmt.Sprintf("(%d, %d, %d)", status.userId, status.checkpointId, status.roadmapId))
	}
	query := fmt.Sprintf("INSERT INTO checkpoint_status (user_id, checkpoint_id, roadmap_id) VALUES %s", strings.Join(values, ", "))

	if _, err := s.DB.Exec(query); err != nil {
		return err
	}
	return nil
}

type DeleteCheckpointStatus struct {
	roadmapId     int
	userIds       []int
	checkpointIds []int
}

func (s *CheckpointStatusStore) DeleteManyCheckpointStatus(input *DeleteCheckpointStatus) error {
	var userIdQuery string
	if len(input.userIds) == 1 {
		userIdQuery = fmt.Sprintf("= %d", input.userIds[0])
	} else {
		strUserIds := make([]string, len(input.userIds))
		for i, id := range input.userIds {
			strUserIds[i] = strconv.Itoa(id)
		}
		userIdQuery = fmt.Sprintf("IN (%s)", strings.Join(strUserIds, ", "))
	}

	var checkpointIdQuery string
	if len(input.checkpointIds) == 1 {
		checkpointIdQuery = fmt.Sprintf("= %d", input.checkpointIds[0])
	} else {
		strCheckpointIds := make([]string, len(input.checkpointIds))
		for i, id := range input.checkpointIds {
			strCheckpointIds[i] = strconv.Itoa(id)
		}
		checkpointIdQuery = fmt.Sprintf("IN (%s)", strings.Join(strCheckpointIds, ", "))
	}

	query := fmt.Sprintf("DELETE FROM checkpoint_status WHERE roadmap_id = %d AND user_id %s AND checkpoint_id %s", input.roadmapId, userIdQuery, checkpointIdQuery)
	if _, err := s.DB.Exec(query); err != nil {
		return err
	}
	return nil

}
