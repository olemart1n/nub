// Package router handles http requests
package router

import (
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"html/template"
	"net/http"
)

func NewRouter(pool *pgxpool.Pool, tmpl *template.Template) *mux.Router {
	r := mux.NewRouter()
	r.Handle("/", Index(tmpl)).Methods("GET")

	r.Handle("/get-upload-form", GetUploadForm(tmpl)).Methods("GET")
	r.Handle("/get-uploader", GetUploader(tmpl)).Methods("GET")

	// Serve static files at /static/
	staticFileDirectory := http.Dir("static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET", "HEAD")
	return r
}
