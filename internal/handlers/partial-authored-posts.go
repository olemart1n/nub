package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/olemart1n/nub/internal/db"
	contextkeys "github.com/olemart1n/nub/internal/handlers/context-keys"
)

func PartialAuthoredPosts(tpl *template.Template, DB *db.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		userID := r.Context().Value(contextkeys.UserIDKey).(string)
		userIDInt, _ := strconv.Atoi(userID)

		posts, err := DB.GetPostsByUserID(r.Context(), userIDInt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tpl.ExecuteTemplate(w, "profile-page-posts", posts)
		if err != nil {
			fmt.Print(err)
			http.Error(w, "error when executing comments.html", http.StatusInternalServerError)
		}

	}
}
