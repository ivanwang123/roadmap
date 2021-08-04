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
	checkpoint_postgres "github.com/ivanwang123/roadmap/internal/checkpoint/repository/postgres"
	checkpoint_usecase "github.com/ivanwang123/roadmap/internal/checkpoint/usecase"
	checkpoint_status_postgres "github.com/ivanwang123/roadmap/internal/checkpoint_status/repository/postgres"
	checkpoint_status_usecase "github.com/ivanwang123/roadmap/internal/checkpoint_status/usecase"
	"github.com/ivanwang123/roadmap/internal/common/auth"
	"github.com/ivanwang123/roadmap/internal/common/cookie"
	"github.com/ivanwang123/roadmap/internal/loaders"
	roadmap_postgres "github.com/ivanwang123/roadmap/internal/roadmap/repository/postgres"
	roadmap_usecase "github.com/ivanwang123/roadmap/internal/roadmap/usecase"
	roadmap_follower_postgres "github.com/ivanwang123/roadmap/internal/roadmap_follower/repository/postgres"
	roadmap_follower_usecase "github.com/ivanwang123/roadmap/internal/roadmap_follower/usecase"
	user_postgres "github.com/ivanwang123/roadmap/internal/user/repository/postgres"
	user_usecase "github.com/ivanwang123/roadmap/internal/user/usecase"
	"github.com/ivanwang123/roadmap/resolvers"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const defaultPort = "8080"

func main() {
	db, err := database.ConnectDB("postgres://postgres:postgres@localhost/roadmap?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	// TODO: Move to resolver, create NewResolver function
	checkpointRepo := checkpoint_postgres.NewCheckpointRepo(db)
	checkpointStatusRepo := checkpoint_status_postgres.NewCheckpointStatusRepo(db)
	roadmapRepo := roadmap_postgres.NewRoadmapRepo(db)
	roadmapFollowerRepo := roadmap_follower_postgres.NewRoadmapFollowerRepo(db)
	userRepo := user_postgres.NewUserRepo(db)

	checkpointUsecase := checkpoint_usecase.NewCheckpointUsecase(checkpointRepo, checkpointStatusRepo, roadmapFollowerRepo)
	checkpointStatusUsecase := checkpoint_status_usecase.NewCheckpointStatusUsecase(checkpointStatusRepo, checkpointRepo)
	roadmapUsecase := roadmap_usecase.NewRoadmapUsecase(roadmapRepo)
	roadmapFollowerUsecase := roadmap_follower_usecase.NewRoadmapFollowerUsecase(roadmapFollowerRepo, roadmapRepo, checkpointRepo, checkpointStatusRepo)
	userUsecase := user_usecase.NewUserUsecase(userRepo)

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
	// router.Use(stores.Middleware(store))
	router.Use(loaders.Middleware(db))

	// TODO: Add IsUnAuthenticated directive
	c := generated.Config{
		Resolvers: &resolvers.Resolver{
			CheckpointUsecase:       checkpointUsecase,
			CheckpointStatusUsecase: checkpointStatusUsecase,
			RoadmapUsecase:          roadmapUsecase,
			RoadmapFollowerUsecase:  roadmapFollowerUsecase,
			UserUsecase:             userUsecase,
		},
	}
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
