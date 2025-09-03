package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/olemart1n/nub/internal/db"
)

func DeletePost(db *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		postIDStr := mux.Vars(r)["id"]
		postID, _ := strconv.Atoi(postIDStr)
		err := db.DeletePost(r.Context(), postID)
		if err != nil {
			http.Error(w, "Failed to delete post: $s"+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", " text/plain")
		w.Write([]byte("Post deleted"))
	}

}
