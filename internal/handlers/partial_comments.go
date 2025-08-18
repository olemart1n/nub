package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/olemart1n/nub/internal/db"
)

func PartialComments(DB *db.DB, tpl *template.Template) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		postIDStr := mux.Vars(r)["id"]
		postID, _ := strconv.Atoi(postIDStr)

		comments, err := DB.GetPostComments(r.Context(), postID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tpl.ExecuteTemplate(w, "comments.html", comments)
		if err != nil {
			http.Error(w, "error when executing comments.html", http.StatusInternalServerError)
		}

	}
}
