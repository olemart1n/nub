// Package middleware handles authentication logic
package middleware

import (
	"context"
	"net/http"

	contextkeys "github.com/olemart1n/nub/internal/handlers/context-keys"
	"github.com/olemart1n/nub/internal/session"
)

func Authenticate(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("session_id")
		if err != nil {
			//		http.Error(w, "Unauthorized", http.StatusUnauthorized)

			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		userID, err := session.GetSession(r.Context(), cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			//	http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), contextkeys.UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
