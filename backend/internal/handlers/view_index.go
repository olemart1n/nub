package handlers

import (
	"html/template"
	"log"
	"net/http"

	contextkeys "github.com/olemart1n/nub/internal/handlers/context-keys"
)

func ViewIndex(tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := TemplateDataIndex{
			Title:      "Nub Global Index",
			UserID:     r.Context().Value(contextkeys.UserIDKey).(string),
			IsLoggedIn: r.Context().Value(contextkeys.IsLoggedInKey).(bool),
			Query:      r.URL.Query().Get("q"),
			Keywords:   Keywords,
		}

		err := tpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			log.Println("Template execution error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
