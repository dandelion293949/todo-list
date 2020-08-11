//go:generate go run github.com/99designs/gqlgen
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
  "github.com/rs/cors"
	"github.com/dandelion293949/todo-list/graph"
	"github.com/dandelion293949/todo-list/graph/domains"
	"github.com/dandelion293949/todo-list/graph/generated"
)

const defaultPort = "9080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

  srv := handler.NewDefaultServer(
    generated.NewExecutableSchema(
      generated.Config{
        Resolvers: &graph.Resolver{
          TodoRepository: domains.NewTodoRepository(),
        },
      },
    ),
  )

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

  c := cors.New(cors.Options {
    AllowedOrigins: []string{"http://localhost:8080",},
    AllowCredentials: true,
  })

	http.Handle("/query", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
