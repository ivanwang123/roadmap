package tests

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/ivanwang123/roadmap/server/database"
	"github.com/ivanwang123/roadmap/server/stores"
)

var (
	TestStore *stores.Store
	Fixtures  *testfixtures.Loader
)

func TestMain(m *testing.M) {
	db, err := sql.Open("pgx", "dbname=roadmap_test password=postgres")
	if err != nil {
		log.Fatalf("Failed to open test database: %s", err)
	}

	Fixtures, err = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgresql"),
		testfixtures.Directory("./fixtures"),
	)
	if err != nil {
		log.Fatalf("Failed to create fixtures: %s", err)
	}

	sqlxDB := database.ConvertSqlDB(db)
	userStore := &stores.UserStore{
		DB: sqlxDB,
	}
	roadmapStore := &stores.RoadmapStore{
		DB: sqlxDB,
	}
	roadmapFollowerStore := &stores.RoadmapFollowerStore{
		DB: sqlxDB,
	}
	checkpointStore := &stores.CheckpointStore{
		DB: sqlxDB,
	}
	checkpointStatusStore := &stores.CheckpointStatusStore{
		DB: sqlxDB,
	}
	TestStore = &stores.Store{
		UserStore:             userStore,
		RoadmapStore:          roadmapStore,
		RoadmapFollowerStore:  roadmapFollowerStore,
		CheckpointStore:       checkpointStore,
		CheckpointStatusStore: checkpointStatusStore,
	}

	os.Exit(m.Run())
}

func prepareTestDatabase() {
	if err := Fixtures.Load(); err != nil {
		log.Fatalf("Failed to load fixtures: %s", err)
	}
}
