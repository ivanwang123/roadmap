package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/ivanwang123/roadmap/database"
	"github.com/ivanwang123/roadmap/graphql/generated"
	"github.com/ivanwang123/roadmap/internal/common/auth"
	"github.com/ivanwang123/roadmap/internal/common/cookie"
	"github.com/ivanwang123/roadmap/internal/loaders"
	"github.com/ivanwang123/roadmap/internal/stores"
	"github.com/ivanwang123/roadmap/resolvers"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const defaultPort = "8080"

func main() {
	db, err := database.ConnectDB("postgres://postgres:postgres@localhost/roadmap?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	store := stores.NewStore(db)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://roadmapper.vercel.app"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Origin", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	router.Use(cookie.Middleware())
	router.Use(auth.Middleware())
	router.Use(database.Middleware(db))
	router.Use(stores.Middleware(store))
	router.Use(loaders.Middleware(db))

	// TODO: Add IsUnAuthenticated directive
	c := generated.Config{Resolvers: &resolvers.Resolver{}}
	c.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		userId := auth.ForContext(ctx)
		fmt.Println("IS AUTHENTICATED", userId)
		if userId > 0 {
			return next(ctx)
		} else {
			return nil, errors.New("Access denied")
		}
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
