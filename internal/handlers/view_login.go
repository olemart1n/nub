package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

func ViewLogin(views *jet.Set) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := views.GetTemplate("pages/login.jet")
		if err != nil {
			log.Println("Template load error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		vars := make(jet.VarMap)
		vars.Set("title", "Signup")

		err = tmpl.Execute(w, vars, nil)
		if err != nil {
			log.Println("Template execution error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
