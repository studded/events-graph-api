package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"

	"github.com/studded/events-graph-api/graph/model"
)

// Roles is the resolver for the roles field.
func (r *userResolver) Roles(ctx context.Context, obj *model.User) ([]*model.Role, error) {
	return r.RolesRepo.GetRolesByUserID(obj.ID)
}

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
