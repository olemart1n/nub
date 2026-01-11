package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/olemart1n/nub/internal/db"
)

func GetPosts(DB *db.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		pageNrStr := mux.Vars(r)["pageNumber"]
		pageNumber, _ := strconv.Atoi(pageNrStr)
		posts, err := DB.GetPostsWithImg(r.Context(), pageNumber)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(posts); err != nil {
			http.Error(w, "go server failed to encode response", http.StatusInternalServerError)
		}

	}
}
