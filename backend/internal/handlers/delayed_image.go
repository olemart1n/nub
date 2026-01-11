package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/olemart1n/nub/internal/db"
)

func DelayedImage(DB *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		imageURL, err := DB.GetImageWithDelay(r.Context())
		if err != nil {
			fmt.Print(err)
			http.Error(w, "error in select query", http.StatusInternalServerError)
			return
		}

		// Add 2 second delay
		time.Sleep(2 * time.Second)

		// Set response headers
		w.Header().Set("Content-Type", "application/json")

		// Create response object
		response := map[string]string{"imageUrl": imageURL}

		// Encode and send response
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	}
}
