package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/olemart1n/nub/internal/db"
	contextkeys "github.com/olemart1n/nub/internal/handlers/context-keys"
)

func ViewProfile(tpl *template.Template, DB *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(contextkeys.UserIDKey).(string)

		userData, err := DB.GetUserPrivate(r.Context(), userID)
		if err != nil {
			fmt.Print(err.Error())
			http.Error(w, "User id not found in the request context or some other error occured", http.StatusUnauthorized)
			return
		}
		var email string
		if userData.Email.Valid {
			email = userData.Email.String
		} else {
			email = "no email"
		}
		data := TemplateDataProfile{
			Title:      "My Profile",
			UserID:     userID,
			Email:      email,
			Username:   userData.Username,
			IsLoggedIn: r.Context().Value(contextkeys.IsLoggedInKey).(bool),
		}
		fmt.Print(data.Username)

		err = tpl.ExecuteTemplate(w, "profile.html", data)
		if err != nil {
			log.Println("Template execution error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
