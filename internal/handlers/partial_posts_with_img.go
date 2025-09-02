package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/olemart1n/nub/internal/db"
)

func PartialPostsWithImg(DB *db.DB, tpl *template.Template) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		pageNrStr := mux.Vars(r)["pageNumber"]
		pageNumber, _ := strconv.Atoi(pageNrStr)
		posts, err := DB.GetPostsWithImg(r.Context(), pageNumber)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tpl.ExecuteTemplate(w, "posts-with-img", posts)
		if err != nil {
			http.Error(w, "error when executing images.html", http.StatusInternalServerError)
		}

	}
}
