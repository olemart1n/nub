package handlers

import (
	"html/template"
	"net/http"

	"github.com/olemart1n/nub/internal/db"
)

func PartialLatestImgs(DB *db.DB, tpl *template.Template, page int) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		images, err := DB.GetLatestImages(r.Context(), page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tpl.ExecuteTemplate(w, "images.html", images)
		if err != nil {
			http.Error(w, "error when executing images.html", http.StatusInternalServerError)
		}

	}
}
