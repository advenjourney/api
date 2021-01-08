package main

import (
	"log"
	"net/http"
	"os"

	"github.com/advenjourney/api/graph"
	"github.com/advenjourney/api/graph/generated"
	"github.com/advenjourney/api/internal/auth"
	_ "github.com/advenjourney/api/internal/auth"
	database "github.com/advenjourney/api/internal/pkg/db/mysql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(auth.Middleware())

	database.InitDB()
	database.Migrate()
	server := handler.GraphQL(api.NewExecutableSchema(api.Config{Resolvers: &api.Resolver{}}))
	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
