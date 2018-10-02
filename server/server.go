package main

import (
	"log"
	"net/http"
	"os"

	"prisma-go-hackernews/prisma-client"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "4000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	client := prisma.New(nil)
	resolver := Resolver{
		Prisma: &client,
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(NewExecutableSchema(Config{Resolvers: &resolver})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
