package tests

import (
	"database/sql"
	"testing"

	"github.com/ivanwang123/roadmap/models"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/stretchr/testify/assert"
)

func TestToggleFollowRoadmap(t *testing.T) {
	prepareTestDatabase()

	var err error
	userId := 2
	roadmapId := 1

	t.Run("Follow roadmap", func(t *testing.T) {
		_, err = TestStore.RoadmapFollowerStore.ToggleFollowRoadmap(userId, &models.FollowRoadmap{RoadmapID: roadmapId})
		assert.Nilf(t, err, "Follow roadmap")

		roadmapFollower, err := TestStore.RoadmapFollowerStore.Get(userId, roadmapId)
		assert.Nilf(t, err, "Check if roadmap follower was created successfully")
		assert.NotNilf(t, roadmapFollower, "Check if roadmap follower was created successfully")

		_, err = TestStore.CheckpointStore.GetCheckpointStatus(userId, 1, roadmapId)
		assert.Nilf(t, err, "Check if checkpoint status 1 was created successfully")
		_, err = TestStore.CheckpointStore.GetCheckpointStatus(userId, 2, roadmapId)
		assert.Nilf(t, err, "Check if checkpoint status 2 was created successfully")
	})

	t.Run("Unfollow roadmap", func(t *testing.T) {
		_, err := TestStore.RoadmapFollowerStore.ToggleFollowRoadmap(userId, &models.FollowRoadmap{RoadmapID: roadmapId})
		assert.Nilf(t, err, "Unfollow roadmap")

		_, err = TestStore.RoadmapFollowerStore.Get(userId, roadmapId)
		assert.ErrorIsf(t, err, sql.ErrNoRows, "Check if roadmap follower was deleted")

		_, err = TestStore.CheckpointStore.GetCheckpointStatus(userId, 1, roadmapId)
		assert.ErrorIsf(t, err, sql.ErrNoRows, "Check if checkpoint status 1 was deleted")
		_, err = TestStore.CheckpointStore.GetCheckpointStatus(userId, 2, roadmapId)
		assert.ErrorIsf(t, err, sql.ErrNoRows, "Check if checkpoint status 2 was deleted")
	})
}
