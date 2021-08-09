package usecase_test

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/ivanwang123/roadmap/models"
	"github.com/ivanwang123/roadmap/resolvers"
	"github.com/ivanwang123/roadmap/tests"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

var (
	resolver *resolvers.Resolver
	fixtures *testfixtures.Loader
	db       *sqlx.DB
)

func TestMain(m *testing.M) {
	resolver, fixtures, db = tests.Setup()

	os.Exit(m.Run())
}

func TestToggleFollowRoadmap(t *testing.T) {
	tests.PrepareTestDatabase(fixtures)

	userID := 2
	roadmapID := 1

	t.Run("Follow roadmap", func(t *testing.T) {
		ctx := context.Background()
		_, err := resolver.RoadmapUsecase.ToggleFollow(ctx, userID, roadmapID)
		assert.Nilf(t, err, "Error following roadmap: %s", err)

		roadmapFollower, err := resolver.RoadmapFollowerUsecase.Get(ctx, userID, roadmapID)
		assert.Nilf(t, err, "Error getting roadmap follower: %s", err)
		assert.NotNilf(t, roadmapFollower, "Roadmap follower was not created")

		_, err = resolver.CheckpointStatusUsecase.Get(ctx, &models.GetCheckpointStatus{UserID: userID, RoadmapID: roadmapID, CheckpointID: 1})
		assert.Nilf(t, err, "Error getting checkpoint status 1: %s", err)
		_, err = resolver.CheckpointStatusUsecase.Get(ctx, &models.GetCheckpointStatus{UserID: userID, RoadmapID: roadmapID, CheckpointID: 2})
		assert.Nilf(t, err, "Error getting checkpoint status 2: %s", err)
	})

	t.Run("Unfollow roadmap", func(t *testing.T) {
		ctx := context.Background()
		_, err := resolver.RoadmapUsecase.ToggleFollow(ctx, userID, roadmapID)
		assert.Nilf(t, err, "Error unfollowing roadmap: %s", err)

		_, err = resolver.RoadmapFollowerUsecase.Get(ctx, userID, roadmapID)
		assert.ErrorIsf(t, err, sql.ErrNoRows, "Roadmap follower was not deleted: %s", err)

		_, err = resolver.CheckpointStatusUsecase.Get(ctx, &models.GetCheckpointStatus{UserID: userID, RoadmapID: roadmapID, CheckpointID: 1})
		assert.ErrorIsf(t, err, sql.ErrNoRows, "Checkpoint status 1 was not deleted: %s", err)
		_, err = resolver.CheckpointStatusUsecase.Get(ctx, &models.GetCheckpointStatus{UserID: userID, RoadmapID: roadmapID, CheckpointID: 2})
		assert.ErrorIsf(t, err, sql.ErrNoRows, "Checkpoint status 2 was not deleted: %s", err)
	})
}
