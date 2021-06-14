package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/ivanwang123/roadmap/server/graph/generated"
	"github.com/ivanwang123/roadmap/server/graph/resolvers"
	"github.com/ivanwang123/roadmap/server/middleware"
	"github.com/ivanwang123/roadmap/server/stores"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const defaultPort = "8080"

// var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
// var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// func toSnakeCase(str string) string {
// 	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
// 	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
// 	return strings.ToLower(snake)
// }

// func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
// 	for _, model := range b.Models {
// 		for _, field := range model.Fields {
// 			// if model.Name == "User" && field.Name == "password" {
// 			// 	field.Tag = fmt.Sprintf(`json:"-" validate:"omitempty"`)
// 			// } else {
// 			field.Tag = fmt.Sprintf(`json:"%s" sql:"%s"`, toSnakeCase(field.Name), toSnakeCase(field.Name))
// 			// }
// 		}
// 	}

// 	return b
// }

func main() {
	// cfg, err := config.LoadConfigFromDefaultLocations()
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
	// 	os.Exit(2)
	// }

	// // Attaching the mutation function onto modelgen plugin
	// p := modelgen.Plugin{
	// 	MutateHook: mutateHook,
	// }

	// err = api.Generate(cfg,
	// 	api.NoPlugins(),
	// 	api.AddPlugin(&p),
	// )
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err.Error())
	// 	os.Exit(3)
	// }

	store, err := stores.NewStore("postgres://postgres:postgres@localhost/roadmap?sslmode=disable")
	if err != nil {
		log.Fatal("Error creating store:", err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(middleware.StoreMiddleware(store))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
