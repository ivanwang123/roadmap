package usecase_test

import (
	"context"
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

func TestAuthenticate(t *testing.T) {
	tests.PrepareTestDatabase(fixtures)

	ctx := context.Background()

	username := "test"
	email := "test@test.com"
	password := "password"

	t.Run("Authenticate with email", func(t *testing.T) {
		user, err := resolver.UserUsecase.Authenticate(ctx, &models.Login{Email: &email, Password: password})
		assert.Nilf(t, err, "Error authenticating user: %s", err)
		assert.NotNilf(t, user, "User was not retrieved")
	})

	t.Run("Authenticate with username", func(t *testing.T) {
		user, err := resolver.UserUsecase.Authenticate(ctx, &models.Login{Username: &username, Password: password})
		assert.Nilf(t, err, "Error authenticating user: %s", err)
		assert.NotNilf(t, user, "User was not retrieved")
	})

	t.Run("Missing credentials", func(t *testing.T) {
		user, err := resolver.UserUsecase.Authenticate(ctx, &models.Login{Password: password})
		assert.NotNilf(t, err, "Expected error")
		assert.Nilf(t, user, "Expected no user, but got: %s", user)
	})

	t.Run("Incorrect credentials", func(t *testing.T) {
		invalidUsername := "invalid"
		user, err := resolver.UserUsecase.Authenticate(ctx, &models.Login{Username: &invalidUsername, Password: password})
		assert.NotNilf(t, err, "Expected error")
		assert.Nilf(t, user, "Expected no user, but got: %s", user)

		invalidPassword := "invalid"
		user, err = resolver.UserUsecase.Authenticate(ctx, &models.Login{Username: &username, Password: invalidPassword})
		assert.NotNilf(t, err, "Expected error")
		assert.Nilf(t, user, "Expected no user, but got: %s", user)
	})
}
