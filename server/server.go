package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/ivanwang123/roadmap/server/auth"
	"github.com/ivanwang123/roadmap/server/dataloaders"
	"github.com/ivanwang123/roadmap/server/graph/generated"
	"github.com/ivanwang123/roadmap/server/graph/resolvers"
	"github.com/ivanwang123/roadmap/server/stores"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

const defaultPort = "8080"

func main() {
	db, err := ConnectDB("postgres://postgres:postgres@localhost/roadmap?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	store := stores.NewStore(db)

	router := chi.NewRouter()

	router.Use(stores.Middleware(store))
	router.Use(dataloaders.Middleware(db))
	router.Use(auth.Middleware(store))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

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
