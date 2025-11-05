package bunny

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/olemart1n/nub/config"
	contextkeys "github.com/olemart1n/nub/internal/handlers/context-keys"
)

func UploadImages(envs config.EnvConfig, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var uploadedURLs []string

		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		files := r.MultipartForm.File["images"]
		if len(files) == 0 {
			http.Error(w, "No files uploaded", http.StatusBadRequest)
			return
		}

		for _, fileHeader := range files {
			file, err := fileHeader.Open()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer file.Close()
			//
			//
			fileBytes, err := io.ReadAll(file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			bunnyURL := fmt.Sprintf("https://%s/%s/%s", envs.StorageRegion, envs.StorageZone, fileHeader.Filename)

			req, err := http.NewRequest("PUT", bunnyURL, bytes.NewReader(fileBytes))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			req.Header.Set("AccessKey", envs.StoragePassword)
			req.Header.Set("Content-Type", "application/octet-stream")

			resp, err := http.DefaultClient.Do(req)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			defer resp.Body.Close()

			if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
				body, _ := io.ReadAll(resp.Body)
				http.Error(w, fmt.Sprintf("Upload failed: %s", string(body)), resp.StatusCode)
				return
			}

			// Public CDN URL
			publicURL := fmt.Sprintf("https://%s/%s", envs.PullZone, fileHeader.Filename)
			uploadedURLs = append(uploadedURLs, publicURL)

		}
		ctx := context.WithValue(r.Context(), contextkeys.ImageURLsKey, uploadedURLs)
		fmt.Println("uploading to bunny successfull")
		next.ServeHTTP(w, r.WithContext(ctx))
	}

}
