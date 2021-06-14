package stores

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/jmoiron/sqlx"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

var storeCtxKey = "store"

type Store struct {
	*UserStore
	*CheckpointStore
	*RoadmapStore
}

func NewStore(dbString string) (*Store, error) {
	db, err := sqlx.Connect("pgx", dbString)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to store: %w", err)
	}

	db.MapperFunc(func(s string) string {
		return toSnakeCase(s)
	})

	return &Store{
		UserStore:       &UserStore{DB: db},
		CheckpointStore: &CheckpointStore{DB: db},
		RoadmapStore:    &RoadmapStore{DB: db},
	}, nil
}

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func Middleware(store *Store) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), storeCtxKey, store)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *Store {
	raw, _ := ctx.Value(storeCtxKey).(*Store)
	return raw
}
