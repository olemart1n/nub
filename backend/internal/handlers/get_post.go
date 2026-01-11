package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/olemart1n/nub/internal/db"
)

func GetPost(tpl *template.Template, DB *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var data Post
		params := mux.Vars(r)
		id := params["id"]
		postID, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "postID not found", http.StatusInternalServerError)
			return
		}

		// FETCH POST FROM DB
		post, err := DB.GetPost(r.Context(), postID)
		if err != nil {
			fmt.Print(err)
			http.Error(w, "error in select query", http.StatusInternalServerError)
			return
		}
		// FETCH IMAGE DATA FROM DB
		images, err := DB.GetImagesByPostID(r.Context(), post.ID)
		if err != nil {
			fmt.Print(err)
			http.Error(w, "error in when selecting images", http.StatusInternalServerError)
			return
		}

		data.Post.UserID = post.UserID
		data.Images = images
		data.Post = post

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "go server failed to encode response", http.StatusInternalServerError)
		}

	}
}
