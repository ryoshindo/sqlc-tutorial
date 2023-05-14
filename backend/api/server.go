package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"github.com/ryoshindo/sqlc-tutorial/backend/api/graph"
	"github.com/ryoshindo/sqlc-tutorial/backend/api/graph/generated"
	"github.com/ryoshindo/sqlc-tutorial/backend/api/graph/session"
	"github.com/ryoshindo/sqlc-tutorial/backend/db"
	repo "github.com/ryoshindo/sqlc-tutorial/backend/repository/sqlc"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := chi.NewRouter()

	d, err := sql.Open("postgres", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatalln(err)
	}

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"}, // FIXME
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	r.Use(middleware.Logger)
	r.Use(newSessionMiddleware())
	r.Use(dbContextMiddleware(d))

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
	return graph.NewResolver(
		repo.NewAccountRepository(),
		repo.NewSessionRepository(),
	)
}

func newSessionMiddleware() func(http.Handler) http.Handler {
	return session.Middleware(repo.NewAccountRepository(), repo.NewSessionRepository())
}

func dbContextMiddleware(d *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = db.Set(ctx, d)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
