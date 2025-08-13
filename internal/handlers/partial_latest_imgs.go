package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/olemart1n/nub/internal/db"
)

func PartialLatestImgs(DB *db.DB, views *jet.Set, page int) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		vars := make(jet.VarMap)

		tmpl, err := views.GetTemplate("partials/image.jet")
		if err != nil {
			fmt.Print(err)
			return
		}

		images, err := DB.GetLatestImages(r.Context(), page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		vars.Set("images", images)

		err = tmpl.Execute(w, vars, nil)
		if err != nil {
			log.Println("Template execution error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//		err = tmpl.Execute(w, vars, nil)

	}
}
