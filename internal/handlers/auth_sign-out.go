package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/olemart1n/nub/internal/db"
	contextkeys "github.com/olemart1n/nub/internal/handlers/context-keys"
	"github.com/olemart1n/nub/internal/session"
)

func AuthSignOut(db *db.DB, tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID := r.Context().Value(contextkeys.UserIDKey).(string)
		sessionErr := session.DeleteSession(r.Context(), userID)

		if sessionErr != nil {
			err := tpl.ExecuteTemplate(w, "notification", Notification{Error: true, Message: sessionErr.Error()})
			if err != nil {
				fmt.Print(err)
				http.Error(w, "error when executing notifiation", http.StatusInternalServerError)
			}
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			Expires:  time.Unix(0, 0), // Expire immediately
			MaxAge:   -1,              // Also force delete
		})

		//http.Redirect(w, r, "/", http.StatusSeeOther)
		w.Header().Set("HX-Redirect", "/")
	}
}
