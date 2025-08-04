package router

import (
	"html/template"
	"net/http"
)

func GetUploadForm(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tmpl.ExecuteTemplate(w, "upload-form", nil)
	}
}
