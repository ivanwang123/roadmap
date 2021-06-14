package dataloaders

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
)

const loaderCtxKey = "dataloader"

type Loader struct {
	UserById *UserLoader
}

func Middleware(db *sqlx.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), loaderCtxKey, &Loader{
				UserById: UserDataloader(db),
			})

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *Loader {
	raw, _ := ctx.Value(loaderCtxKey).(*Loader)
	return raw
}
