package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/studded/events-graph-api/graph"
	customMiddleware "github.com/studded/events-graph-api/middleware"
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

	userRepo := postgres.UsersRepo{DB: db}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	// just assign user for checking permissions (auth would go here)
	router.Use(customMiddleware.AssignCurrentUser(userRepo))

	config := graph.Config{Resolvers: &graph.Resolver{
		ActivitiesRepo: postgres.ActivitiesRepo{DB: db},
		EventsRepo:     postgres.EventsRepo{DB: db},
		ExpensesRepo:   postgres.ExpensesRepo{DB: db},
		RolesRepo:      postgres.RolesRepo{DB: db},
		UsersRepo:      userRepo,
	}}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(config))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
