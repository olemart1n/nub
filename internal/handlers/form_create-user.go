package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/olemart1n/nub/internal/db"
)

func FormCreateUser(DB *db.DB, tpl *template.Template) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")
		var notification Notification
		err := DB.CreateUser(r.Context(), username, password)

		if err != nil {
			notification = Notification{
				Error:   true,
				Message: "Something went wrong!",
			}

			err = tpl.ExecuteTemplate(w, "notification.html", notification)
			if err != nil {
				log.Println("Template execution error:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			return
		}

		notification = Notification{
			Error:   false,
			Message: "User " + username + " successfully created",
		}
		err = tpl.ExecuteTemplate(w, "notification.html", notification)
		if err != nil {
			log.Println("Template execution error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
