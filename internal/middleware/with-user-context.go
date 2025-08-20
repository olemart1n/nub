package middleware

import (
	"context"
	"fmt"
	"net/http"

	contextkeys "github.com/olemart1n/nub/internal/handlers/context-keys"
	"github.com/olemart1n/nub/internal/session"
)

// Soft auth: does not redirect, just sets context

func WithUserContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userID string
		cookie, err := r.Cookie("session_id")
		if err == nil {
			if id, err := session.GetSession(r.Context(), cookie.Value); err != nil {
				fmt.Print(err)
			} else {

				userID = id
			}

		}

		ctx := context.WithValue(r.Context(), contextkeys.UserIDKey, userID)
		ctx = context.WithValue(ctx, contextkeys.IsLoggedInKey, userID != "")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
