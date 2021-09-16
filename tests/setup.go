package tests

import (
	"database/sql"
	"log"
	"path/filepath"
	"runtime"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/ivanwang123/roadmap/database"
	checkpoint_usecase "github.com/ivanwang123/roadmap/internal/checkpoint/usecase"
	checkpoint_status_usecase "github.com/ivanwang123/roadmap/internal/checkpoint_status/usecase"
	roadmap_usecase "github.com/ivanwang123/roadmap/internal/roadmap/usecase"
	roadmap_follower_usecase "github.com/ivanwang123/roadmap/internal/roadmap_follower/usecase"
	user_usecase "github.com/ivanwang123/roadmap/internal/user/usecase"
	"github.com/ivanwang123/roadmap/resolvers"
	"github.com/ivanwang123/roadmap/tests/mocks"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
)

func Setup() (*resolvers.Resolver, *testfixtures.Loader, *sqlx.DB) {
	// TODO: Use roadmap_test for ci/cd test
	db, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/roadmap_test?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to open test database: %s", err)
	}
	// TODO: Remove?
	// db.SetMaxOpenConns(1)

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	fixturepath := filepath.Join(basepath, "fixtures")

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgresql"),
		testfixtures.Directory(fixturepath),
	)
	if err != nil {
		log.Fatalf("Failed to create fixtures: %s", err)
	}

	sqlxDB := database.ConvertSqlDB(db)
	resolver := resolvers.NewResolver(sqlxDB)

	return resolver, fixtures, sqlxDB
}

func PrepareTestDatabase(fixtures *testfixtures.Loader) {
	if err := fixtures.Load(); err != nil {
		log.Fatalf("Failed to load fixtures: %s", err)
	}
}

func CreateMockResolver() *resolvers.Resolver {
	checkpointRepo := new(mocks.CheckpointMockRepo)
	checkpointStatusRepo := new(mocks.CheckpointStatusMockRepo)
	roadmapRepo := new(mocks.RoadmapMockRepo)
	roadmapFollowerRepo := new(mocks.RoadmapFollowerMockRepo)
	userRepo := new(mocks.UserMockRepo)

	checkpointStatusRepo.On("Create", mock.Anything).Return(nil, nil)

	checkpointUsecase := checkpoint_usecase.NewCheckpointUsecase(checkpointRepo, checkpointStatusRepo, roadmapFollowerRepo)
	checkpointStatusUsecase := checkpoint_status_usecase.NewCheckpointStatusUsecase(checkpointStatusRepo, checkpointRepo)
	roadmapUsecase := roadmap_usecase.NewRoadmapUsecase(roadmapRepo, roadmapFollowerRepo, checkpointRepo, checkpointStatusRepo)
	roadmapFollowerUsecase := roadmap_follower_usecase.NewRoadmapFollowerUsecase(roadmapFollowerRepo)
	userUsecase := user_usecase.NewUserUsecase(userRepo)

	return &resolvers.Resolver{
		CheckpointUsecase:       checkpointUsecase,
		CheckpointStatusUsecase: checkpointStatusUsecase,
		RoadmapUsecase:          roadmapUsecase,
		RoadmapFollowerUsecase:  roadmapFollowerUsecase,
		UserUsecase:             userUsecase,
	}
}
