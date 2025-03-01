package handlers

// import (
// 	"context"
// 	"net/http"

// 	"edu/internal/models"
// )

// type contextKey string

// const userContextKey contextKey = "user"

// func (h *Handler) ContextSetUser(r *http.Request, user *models.User) *http.Request {
// 	ctx := context.WithValue(r.Context(), userContextKey, user)
// 	return r.WithContext(ctx)
// }

// func (h *Handler) ContextGetUser(r *http.Request) *models.User {
// 	user, ok := r.Context().Value(userContextKey).(*models.User)
// 	if !ok {
// 		panic("missing user value in request context")
// 	}
// 	return user
// }