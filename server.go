package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/studded/events-graph-api/graph"
	"github.com/studded/events-graph-api/postgres"
)

const defaultPort = "8080"

func main() {
	// connect to postgres repo
	db, err := sqlx.Connect("postgres", "user=studded dbname=krane sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	config := graph.Config{Resolvers: &graph.Resolver{
		ActivitiesRepo: postgres.ActivitiesRepo{DB: db},
		EventsRepo:     postgres.EventsRepo{DB: db},
		RolesRepo:      postgres.RolesRepo{DB: db},
		UsersRepo:      postgres.UsersRepo{DB: db},
	}}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(config))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
