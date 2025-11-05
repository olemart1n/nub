package handlers

import (
	"html/template"
	"net/http"
)

func ViewSignUp(tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := tpl.ExecuteTemplate(w, "sign-up.html", nil)
		if err != nil {
			http.Error(w, "error when executing sign-up.html", http.StatusInternalServerError)
		}

	}
}
