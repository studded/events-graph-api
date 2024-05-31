package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/studded/events-graph-api/graph/model"
	"github.com/studded/events-graph-api/postgres"
)

const CurrentUserKey = "currentUser"

func AssignCurrentUser(userRepo postgres.UsersRepo) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			email := r.Header.Get("User-Email")

			if email == "" {
				next.ServeHTTP(w, r)
				return
			}

			user, err := userRepo.GetUserByEmail(email)

			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), CurrentUserKey, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetCurrentUserFromCTX(ctx context.Context) (*model.User, error) {
	if ctx.Value(CurrentUserKey) == nil {
		return nil, errors.New("no user")
	}

	user, ok := ctx.Value(CurrentUserKey).(*model.User)
	if !ok {
		return nil, errors.New("invalid user")
	}

	return user, nil
}
