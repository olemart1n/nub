// Package middleware handles authentication logic
package middleware

import (
	"context"
	"net/http"

	"github.com/olemart1n/nub/internal/session"
)

func Authenticate(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userID, err := session.GetSession(r.Context(), cookie.Value)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		type contextKey string
		const userIDKey contextKey = "userID"
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
