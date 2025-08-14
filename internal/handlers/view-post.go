package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/mux"
	"github.com/olemart1n/nub/internal/db"
	contextkeys "github.com/olemart1n/nub/internal/handlers/context-keys"
)

func ViewPost(views *jet.Set, DB *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Post db.Post
		var Images []db.Image
		params := mux.Vars(r)
		id := params["id"]
		postID, err := strconv.Atoi(id)
		fmt.Print(postID)
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
		Post = post
		// FETCH IMAGE DATA FROM DB
		images, err := DB.GetImagesByPostID(r.Context(), Post.ID)
		if err != nil {
			fmt.Print(err)
			http.Error(w, "error in when selecting images", http.StatusInternalServerError)
			return
		}
		Images = images

		tmpl, err := views.GetTemplate("pages/post.jet")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		vars := make(jet.VarMap)

		vars.Set("post", Post)
		vars.Set("images", Images)
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
