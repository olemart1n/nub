package handlers

import (
	"html/template"
	"net/http"
)

func ViewSignup(tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var data TemplateDataIndex
		data.Title = "Signup"
		err := tpl.ExecuteTemplate(w, "signup.html", data)
		if err != nil {
			http.Error(w, "error when executing post.html", http.StatusInternalServerError)
		}

	}
}
