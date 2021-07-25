package tests

import (
	"testing"

	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateCheckpoint(t *testing.T) {
	prepareTestDatabase()

	userId := 3
	roadmapId := 1
	checkpointId := -1

	_, err := TestStore.RoadmapFollowerStore.ToggleFollowRoadmap(TestStore, userId, roadmapId)
	assert.Nilf(t, err, "Follow roadmap")

	checkpoint, err := TestStore.CheckpointStore.Create(&model.NewCheckpoint{Title: "New checkpoint", Instructions: "Instructions", Links: []string{}, RoadmapID: roadmapId})
	assert.Nilf(t, err, "Create new checkpoint")
	assert.NotNil(t, checkpoint, "Create new checkpoint")
	checkpointId = checkpoint.ID

	checkpoints, err := TestStore.CheckpointStore.GetByRoadmap(roadmapId)
	assert.Nilf(t, err, "Get checkpoints")
	assert.Lenf(t, checkpoints, 3, "Check if checkpoint was created successfully")

	_, err = TestStore.CheckpointStore.GetCheckpointStatus(userId, checkpointId, roadmapId)
	assert.Nilf(t, err, "Check if checkpoint status was created for existing followers")
}
