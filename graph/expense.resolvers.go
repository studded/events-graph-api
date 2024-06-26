package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"

	"github.com/studded/events-graph-api/graph/model"
)

// Event is the resolver for the event field.
func (r *expenseResolver) Event(ctx context.Context, obj *model.Expense) (*model.Event, error) {
	return r.EventsRepo.GetEventByID(obj.EventID)
}

// Expense returns ExpenseResolver implementation.
func (r *Resolver) Expense() ExpenseResolver { return &expenseResolver{r} }

type expenseResolver struct{ *Resolver }
