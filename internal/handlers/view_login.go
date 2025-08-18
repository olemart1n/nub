package handlers

import (
	"html/template"
	"net/http"
)

func ViewLogin(tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data TemplateDataIndex
		data.Title = "Login"
		err := tpl.ExecuteTemplate(w, "login.html", data)
		if err != nil {
			http.Error(w, "error when executing post.html", http.StatusInternalServerError)
		}

	}
}
