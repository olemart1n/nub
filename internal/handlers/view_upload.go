package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/olemart1n/nub/utils"
)

func ViewUpload(views *jet.Set) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := views.GetTemplate("pages/upload.jet")
		if err != nil {
			log.Println("Template load error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		vars := make(jet.VarMap)
		vars.Set("title", "Upload images to nub.global")
		userID := r.Context()
		vars.Set("userID", userID)
		vars.Set("countries", utils.Countries)
		err = tmpl.Execute(w, vars, nil)
		if err != nil {
			log.Println("Template execution error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
