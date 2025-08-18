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

func FormSubmitComment(DB *db.DB, tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		postIDStr := vars["id"]
		postID, err := strconv.Atoi(postIDStr)
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}
		userID := r.Context().Value(contextkeys.UserIDKey).(string)
		userIDint, _ := strconv.Atoi(string(userID))

		fmt.Print("is this function running?")
		content := r.FormValue("content")

		if len(content) == 0 {
			http.Error(w, "no characters found", http.StatusBadRequest)
			return
		}

		comment, err := DB.CreateComment(r.Context(), userIDint, postID, content)
		if err != nil {
			fmt.Print(err)
			http.Error(w, "Could not create comment", http.StatusInternalServerError)
			return
		}
		fmt.Print("comment: ", comment)
		err = tpl.ExecuteTemplate(w, "comment", comment)
		if err != nil {
			http.Error(w, "error when executing post.html", http.StatusInternalServerError)
		}
	}
}
