package main

import (
	"log"
	"net/http"
	"os"

	"example.com/server-go/graph/generated"
	"example.com/server-go/resolver"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"
const defaultPath = "tvu_graphql"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	gqlPath := os.Getenv("GRAPHQL_PATH")
	if gqlPath == "" {
		gqlPath = defaultPath
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	http.Handle("/tvu_playground", playground.Handler("GraphQL playground", `/${gqlPath}`))
	http.Handle(`/${gqlPath}`, srv)

	log.Printf("connect to http://localhost:%s/tvu_playground for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
