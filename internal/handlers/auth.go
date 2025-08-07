package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/olemart1n/nub/internal/db"
	"github.com/olemart1n/nub/internal/session"
)

func LoginHandler(db *db.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		u, err := db.AuthenticateUser(username, password)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
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
		fmt.Println("Working fine")
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			Expires:  time.Now().Add(24 * time.Hour),
		})

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
