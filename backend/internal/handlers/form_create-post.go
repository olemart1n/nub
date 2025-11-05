package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/olemart1n/nub/internal/db"
	contextkeys "github.com/olemart1n/nub/internal/handlers/context-keys"
)

func FormCreatePost(db *db.DB, tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		imageURLs := r.Context().Value(contextkeys.ImageURLsKey).([]string)

		if imageURLs == nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		userID := r.Context().Value(contextkeys.UserIDKey).(string)
		userIDint, _ := strconv.Atoi(string(userID))

		title := r.FormValue("title")

		var tags []string
		// tagsJSON := r.FormValue("tags")

		// err := json.Unmarshal([]byte(tagsJSON), &tags)
		// if err != nil {
		// 	fmt.Print(err)
		// 	http.Error(w, "Invalid tags format", http.StatusBadRequest)
		// 	return
		// }

		location := r.FormValue("location")

		postID, err := db.CreatePost(r.Context(), userIDint, title, location, tags, imageURLs)
		if err != nil {
			fmt.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			_ = tpl.ExecuteTemplate(w, "notification", Notification{Error: true, Message: err.Error()})
			return
		}

		//notification := Notification{
		//			Error:   false,
		//			Message: "Post successfully created",
		//}

		w.Header().Set("HX-Redirect", "/post/"+strconv.Itoa(postID))
	}
}
