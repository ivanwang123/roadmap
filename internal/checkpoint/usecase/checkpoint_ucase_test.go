package usecase_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	checkpoint_postgres "github.com/ivanwang123/roadmap/internal/checkpoint/repository/postgres"
	"github.com/ivanwang123/roadmap/internal/checkpoint/usecase"
	roadmap_follower_postgres "github.com/ivanwang123/roadmap/internal/roadmap_follower/repository/postgres"
	"github.com/ivanwang123/roadmap/models"
	"github.com/ivanwang123/roadmap/resolvers"
	"github.com/ivanwang123/roadmap/tests"
	"github.com/ivanwang123/roadmap/tests/mocks"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestCreateCheckpoint(t *testing.T) {
	tests.PrepareTestDatabase(fixtures)

	userID := 3
	roadmapID := 1

	ctx := context.Background()

	t.Run("Create checkpoint status for follower", func(t *testing.T) {
		_, err := resolver.RoadmapUsecase.ToggleFollow(ctx, userID, roadmapID)
		assert.Nilf(t, err, "Error following roadmap: %s", err)

		checkpoint, err := resolver.CheckpointUsecase.Create(ctx, &models.NewCheckpoint{Title: "New checkpoint", Instructions: "Instructions", Links: []string{}, RoadmapID: roadmapID})
		assert.Nilf(t, err, "Error creating checkpoint: %s", err)
		assert.NotNil(t, checkpoint, "Checkpoint was not created")
		checkpointID := checkpoint.ID

		checkpoints, err := resolver.CheckpointUsecase.GetByRoadmap(ctx, roadmapID)
		assert.Nilf(t, err, "Error getting checkpoints: %s", err)
		assert.Lenf(t, checkpoints, 3, "Checkpoint was not created")

		_, err = resolver.CheckpointStatusUsecase.Get(ctx, &models.GetCheckpointStatus{UserID: userID, RoadmapID: roadmapID, CheckpointID: checkpointID})
		assert.Nilf(t, err, "Checkpoint status was not created for follower: %s", err)
	})

	t.Run("Checkpoint rollback on error", func(t *testing.T) {
		checkpointRepo := checkpoint_postgres.NewCheckpointRepo(db)
		roadmapFollowerRepo := roadmap_follower_postgres.NewRoadmapFollowerRepo(db)
		checkpointStatusRepo := new(mocks.CheckpointStatusMockRepo)

		checkpointUsecase := usecase.NewCheckpointUsecase(checkpointRepo, checkpointStatusRepo, roadmapFollowerRepo)

		checkpointStatusRepo.On("CreateMany", mock.Anything, mock.Anything).Return(errors.New("Mock error"))

		checkpoints, err := resolver.CheckpointUsecase.GetByRoadmap(ctx, roadmapID)
		assert.Nilf(t, err, "Error getting checkpoints: %s", err)
		numInitialCheckpoints := len(checkpoints)

		checkpoint, err := checkpointUsecase.Create(ctx, &models.NewCheckpoint{Title: "New checkpoint", Instructions: "Instructions", Links: []string{}, RoadmapID: roadmapID})
		assert.Nilf(t, checkpoint, "Checkpoint should not be created")
		assert.NotNil(t, err, "Error should be 'Mock error'")

		checkpoints, err = resolver.CheckpointUsecase.GetByRoadmap(ctx, roadmapID)
		assert.Nilf(t, err, "Error getting checkpoints: %s", err)
		numFinalCheckpoints := len(checkpoints)
		assert.Equalf(t, numInitialCheckpoints, numFinalCheckpoints, "Checkpoint was not rolled back")
	})
}

func TestUpdateStatus(t *testing.T) {
	tests.PrepareTestDatabase(fixtures)

	userID := 3
	roadmapID := 1
	checkpointID := 1

	ctx := context.Background()

	t.Run("Update checkpoint status", func(t *testing.T) {
		_, err := resolver.RoadmapUsecase.ToggleFollow(ctx, userID, roadmapID)
		assert.Nilf(t, err, "Error following roadmap: %s", err)

		initialCheckpointStatus, err := resolver.CheckpointStatusUsecase.Get(ctx, &models.GetCheckpointStatus{UserID: userID, RoadmapID: roadmapID, CheckpointID: checkpointID})
		assert.Nilf(t, err, "Error getting checkpoint status: %s", err)
		assert.Equalf(t, models.StatusIncomplete, initialCheckpointStatus.Status, "Initial checkpoint status is not 'INCOMPLETE'")

		checkpoint, err := resolver.CheckpointUsecase.UpdateStatus(ctx, userID, &models.UpdateStatus{CheckpointID: checkpointID, Status: models.StatusComplete})
		assert.Nilf(t, err, "Error updating status: %s", err)

		assert.Equalf(t, models.StatusType(models.StatusComplete), checkpoint.Status, "Updated checkpoint status is not 'COMPLETE'")
	})
}
