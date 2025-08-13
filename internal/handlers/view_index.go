package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	contextkeys "github.com/olemart1n/nub/internal/handlers/context-keys"
)

func ViewIndex(views *jet.Set) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tmpl, err := views.GetTemplate("pages/index.jet")
		if err != nil {
			log.Println("Template load error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		vars := make(jet.VarMap)

		vars.Set("title", "Nub Global Homepage")

		userID := r.Context().Value(contextkeys.UserIDKey).(string)

		//	userID := r.Context().Value("userID").(string)
		vars.Set("userID", userID)

		err = tmpl.Execute(w, vars, nil)
		if err != nil {
			log.Println("Template execution error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
