package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/studded/events-graph-api/postgres"
)

type Resolver struct {
	ActivitiesRepo postgres.ActivitiesRepo
	EventsRepo     postgres.EventsRepo
	RolesRepo      postgres.RolesRepo
	UsersRepo      postgres.UsersRepo
}
