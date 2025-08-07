// Package bunny handles request for uploading to bunny
package bunny

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/olemart1n/nub/config"
)

func SignHandler(envs config.EnvConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filename := r.URL.Query().Get("filename")
		if filename == "" {
			http.Error(w, "filename required", http.StatusBadRequest)
			return
		}
		expires := time.Now().Add(1 * time.Minute).Unix()
		path := fmt.Sprint("/%s/%s", envs.StorageZone, filename)

		// CREATE SIGNATURE (HMAC SHA256)
		mac := hmac.New(sha256.New, []byte(envs.UploadKey))
		mac.Write([]byte(fmt.Sprintf("%d\n%s", expires, path)))
		signature := hex.EncodeToString(mac.Sum(nil))

		uploadURL := fmt.Sprintf("https://%s%s?Expires=%d&Signature=%s",
			envs.Hostname, path, expires, signature)
		publicURL := fmt.Sprintf("https://%s/%s", envs.PullZone, filename)

		json.NewEncoder(w).Encode(map[string]string{

			"uploadURL": uploadURL,
			"publicURL": publicURL,
		})
		fmt.Println("signHandler function is called")
		w.Header().Set("Content-Type", "application/json")

	}
}
