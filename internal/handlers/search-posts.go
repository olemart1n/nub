package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/olemart1n/nub/internal/db"
)

func SearchPosts(DB *db.DB, tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Grab the search query
		q := r.URL.Query().Get("q")

		pageStr := r.URL.Query().Get("page")
		page := 0
		if p, err := strconv.Atoi(pageStr); err != nil && p >= 0 {
			page = p
		}

		// Retrieve posts with image metadata
		posts, err := DB.SearchPostsWithImg(r.Context(), q, page)

		if err != nil {
			http.Error(w, "Failed to search posts", http.StatusInternalServerError)
			return
		}

		// Render only the posts gallery partial
		err = tpl.ExecuteTemplate(w, "latest-posts-with-img.html", posts)
		if err != nil {
			log.Println("Template execution error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
