// Package bunny handles request for uploading to bunny
package bunny

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/olemart1n/nub/config"
)

func SignHandler(envs config.EnvConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate filename
		filename := r.URL.Query().Get("filename")
		if filename == "" {
			http.Error(w, "filename required", http.StatusBadRequest)
			return
		}
		// Sanitize filename (allow alphanumeric, underscores, hyphens, and common extensions)
		if !regexp.MustCompile(`^[a-zA-Z0-9_-]+\.(jpg|jpeg|png|gif|mp4)$`).MatchString(filename) {
			http.Error(w, "invalid filename", http.StatusBadRequest)
			return
		}

		// Set expiration to 15 minutes
		expires := time.Now().Add(15 * time.Minute).Unix()
		path := fmt.Sprintf("/%s/%s", envs.StorageZone, filename)

		// Create signature (try full path with query parameters)
		mac := hmac.New(sha256.New, []byte(envs.StoragePassword))
		signatureString := fmt.Sprintf("%s?Expires=%d", path, expires)
		mac.Write([]byte(signatureString))
		signature := hex.EncodeToString(mac.Sum(nil))

		// Generate URLs

		log.Printf("Signature string: %s", signatureString)
		uploadURL := fmt.Sprintf("https://%s%s?Expires=%d&Signature=%s",
			envs.StorageRegion, path, expires, signature)
		publicURL := fmt.Sprintf("https://%s/%s", envs.PullZone, filename)

		// Set response headers
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*") // Adjust for your origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

		// Encode response
		if err := json.NewEncoder(w).Encode(map[string]string{
			"uploadURL": uploadURL,
			"publicURL": publicURL,
		}); err != nil {
			log.Printf("Error encoding JSON: %v", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}
}
