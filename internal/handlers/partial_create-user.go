package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/olemart1n/nub/internal/db"
)

func HxCreateUser(DB *db.DB, views *jet.Set) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		tmpl, err := views.GetTemplate("partials/notification.jet")
		vars := make(jet.VarMap)

		if err != nil {
			fmt.Print("GetTemplate error: partials/notification.jet")
		}
		err = DB.CreateUser(r.Context(), username, password)

		if err != nil {
			vars.Set("notification", Notification{
				Error:   true,
				Message: "Something went wrong!",
			})

			err = tmpl.Execute(w, vars, nil)
			if err != nil {
				log.Println("Template load error:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		vars.Set("notification", Notification{
			Error:   false,
			Message: "User " + username + " successfully created",
		})
		err = tmpl.Execute(w, vars, nil)
		if err != nil {
			log.Println("Template execution error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
