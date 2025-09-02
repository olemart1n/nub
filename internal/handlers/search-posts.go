package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/olemart1n/nub/internal/db"
	contextkeys "github.com/olemart1n/nub/internal/handlers/context-keys"
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

		if r.Header.Get("HX-Request") == "true" {
			// Render only the posts gallery partial
			err = tpl.ExecuteTemplate(w, "posts-with-img", posts)
			if err != nil {
				log.Println("Template execution error:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else {

			data := TemplateDataIndex{
				Title:      q,
				UserID:     r.Context().Value(contextkeys.UserIDKey).(string),
				IsLoggedIn: r.Context().Value(contextkeys.IsLoggedInKey).(bool),
				Query:      r.URL.Query().Get("q"),
				Posts:      posts,
			}

			err := tpl.ExecuteTemplate(w, "index.html", data)
			if err != nil {
				log.Println("Template execution error:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
