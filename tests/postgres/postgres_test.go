package postgres_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/ivanwang123/roadmap/database"
	"github.com/ivanwang123/roadmap/resolvers"
)

var (
	TestResolver *resolvers.Resolver
	Fixtures     *testfixtures.Loader
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
	TestResolver = resolvers.NewResolver(sqlxDB)

	os.Exit(m.Run())
}

func PrepareTestDatabase() {
	if err := Fixtures.Load(); err != nil {
		log.Fatalf("Failed to load fixtures: %s", err)
	}
}
