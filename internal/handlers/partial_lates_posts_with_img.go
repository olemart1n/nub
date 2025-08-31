package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/olemart1n/nub/internal/db"
)

func PartialLatestPostsWithImg(DB *db.DB, tpl *template.Template) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		pageNrStr := mux.Vars(r)["pageNumber"]
		pageNumber, _ := strconv.Atoi(pageNrStr)
		fmt.Println("page number: " + pageNrStr)
		posts, err := DB.GetLatestPostsWithImg(r.Context(), pageNumber)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tpl.ExecuteTemplate(w, "latest-posts-with-img.html", posts)
		if err != nil {
			http.Error(w, "error when executing images.html", http.StatusInternalServerError)
		}

	}
}
