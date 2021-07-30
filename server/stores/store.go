package stores

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
)

const storeCtxKey = "store"

type Store struct {
	UserStore             *UserStore
	CheckpointStore       *CheckpointStore
	CheckpointStatusStore *CheckpointStatusStore
	RoadmapStore          *RoadmapStore
	RoadmapFollowerStore  *RoadmapFollowerStore
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		UserStore:             &UserStore{DB: db},
		CheckpointStore:       &CheckpointStore{DB: db},
		CheckpointStatusStore: &CheckpointStatusStore{DB: db},
		RoadmapStore:          &RoadmapStore{DB: db},
		RoadmapFollowerStore:  &RoadmapFollowerStore{DB: db},
	}
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
