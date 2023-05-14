package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/ryoshindo/sqlc-tutorial/backend/api/graph"
	"github.com/ryoshindo/sqlc-tutorial/backend/api/graph/generated"
	repo "github.com/ryoshindo/sqlc-tutorial/backend/repository/sqlc"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := chi.NewRouter()
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"}, // FIXME
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	r.Use(middleware.Logger)

	schemaConfig := generated.Config{Resolvers: newResolver()}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(schemaConfig))

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port) // FIXME: hostname
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalln(err)
	}
}

func newResolver() *graph.Resolver {
	return graph.NewResolver(repo.NewAccountRepository())
}
