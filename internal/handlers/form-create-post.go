package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/olemart1n/nub/internal/db"
	contextkeys "github.com/olemart1n/nub/internal/handlers/context-keys"
)

func FormCreatePost(db *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		imageURLs := r.Context().Value(contextkeys.ImageURLsKey).([]string)

		if imageURLs == nil {
			fmt.Println("imageURLs is indeed nil")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		userID := r.Context().Value(contextkeys.UserIDKey).(string)
		userIDint, _ := strconv.Atoi(string(userID))

		title := r.FormValue("title")
		tagsJSON := r.FormValue("tags")

		if tagsJSON == "" {
			fmt.Println("tagsJSON is empty")
			http.Error(w, "Tags field is required", http.StatusBadRequest)
			return
		}
		var tags []string
		err := json.Unmarshal([]byte(tagsJSON), &tags)
		if err != nil {
			fmt.Print(err)
			http.Error(w, "Invalid tags format", http.StatusBadRequest)
			return
		}
		location := r.FormValue("location")

		err = db.CreatePost(r.Context(), userIDint, title, location, tags, imageURLs)
		if err != nil {
			fmt.Print(err)
			http.Error(w, "Could not create post", http.StatusInternalServerError)
			return
		}

		notification := Notification{
			Error:   false,
			Message: "Post successfully created",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(notification)
	}
}

type PostRequest struct {
	Tags []string `json:"tags"`
}
