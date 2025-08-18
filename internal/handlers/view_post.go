package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/olemart1n/nub/internal/db"
	contextkeys "github.com/olemart1n/nub/internal/handlers/context-keys"
)

func ViewPost(tpl *template.Template, DB *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var data TemplateDataPost
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

		data.Index.UserID = r.Context().Value(contextkeys.UserIDKey).(string)
		data.Index.IsLoggedIn = r.Context().Value(contextkeys.IsLoggedInKey).(bool)
		data.Index.Title = post.Title
		data.Post.UserID = post.UserID
		data.Images = images
		data.Post = post

		err = tpl.ExecuteTemplate(w, "post.html", data)
		if err != nil {
			fmt.Print(err)
			http.Error(w, "error when executing post.html", http.StatusInternalServerError)
		}

	}
}
