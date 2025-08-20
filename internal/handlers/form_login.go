package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/olemart1n/nub/internal/db"
	"github.com/olemart1n/nub/internal/session"
)

func FormLogin(db *db.DB, tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		u, err := db.AuthenticateUser(username, password)
		if err != nil {
			fmt.Println(err)
			err = tpl.ExecuteTemplate(w, "notification", Notification{Error: true, Message: err.Error()})
			if err != nil {
				fmt.Print(err)
				http.Error(w, "error when executing notifiation, can not login", http.StatusInternalServerError)
			}
			return
		}

		userID := strconv.Itoa(u.ID) // Replace with real auth logic
		sessionID := session.GenerateSessionID()
		sessionErr := session.SetSession(r.Context(), sessionID, userID, 24*time.Hour)

		if sessionErr != nil {
			fmt.Println("Redis error")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			Expires:  time.Now().Add(24 * time.Hour),
		})

		//http.Redirect(w, r, "/", http.StatusSeeOther)
		w.Header().Set("HX-Redirect", "/")
	}
}
