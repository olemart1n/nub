package handlers

import (
	"html/template"
	"log"
	"net/http"

	contextkeys "github.com/olemart1n/nub/internal/handlers/context-keys"
)

func ViewIndex(tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var data TemplateDataIndex
		data.Title = "Nub Global Index"

		data.UserID = r.Context().Value(contextkeys.UserIDKey).(string)
		data.IsLoggedIn = r.Context().Value(contextkeys.IsLoggedInKey).(bool)

		err := tpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			log.Println("Template execution error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
