package database

import (
	"context"
	"net/http"
	"regexp"
	"strings"

	"github.com/jmoiron/sqlx"
)

const dbCtxKey = "db"

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func ConnectDB(dbString string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", dbString)
	if err != nil {
		return nil, err
	}

	db.MapperFunc(func(s string) string {
		return toSnakeCase(s)
	})

	return db, nil
}

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func Middleware(db *sqlx.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), dbCtxKey, db)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *sqlx.DB {
	raw, _ := ctx.Value(dbCtxKey).(*sqlx.DB)
	return raw
}
