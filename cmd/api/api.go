package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/advenjourney/api/graph"
	api "github.com/advenjourney/api/graph/generated"
	"github.com/advenjourney/api/internal/auth"
	database "github.com/advenjourney/api/internal/pkg/db/mysql"
	"github.com/advenjourney/api/pkg/config"
	"github.com/advenjourney/api/pkg/version"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	log.Printf("api %s %s", version.Info(), version.BuildContext())

	cfg := config.Load()
	_ = godotenv.Load()
	if addr, ok := os.LookupEnv("API_SERVER_PORT"); ok {
		cfg.Server.Addr = addr
	}
	dsn, ok := os.LookupEnv("API_DB_DSN")
	if !ok || dsn == "" {
		log.Fatal("DSN (API_DB_DSN) not provided")
	}
	cfg.Database.DSN = dsn

	router := chi.NewRouter()
	router.Use(auth.Middleware())

	database.InitDB(*cfg)
	database.Migrate()
	server := handler.NewDefaultServer(api.NewExecutableSchema(api.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("Server running at %s", cfg.Server.Addr)
	log.Fatal(http.ListenAndServe(cfg.Server.Addr, router))
}
