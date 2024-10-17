package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/namanthanki/social/internal/store"
)

type userKey string

const userKeyContext userKey = "user"

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) usersContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		ctx := r.Context()

		post, err := app.store.Users.GetByID(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFoundError(w, r, err)

			default:
				app.internalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, userKeyContext, post)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromContext(r *http.Request) *store.User {
	user, _ := r.Context().Value(userKeyContext).(*store.User)
	return user
}
