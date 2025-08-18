// Package handlers handles http requests
package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/olemart1n/nub/config"
	"github.com/olemart1n/nub/internal/db"
	"github.com/olemart1n/nub/internal/handlers/bunny"
	"github.com/olemart1n/nub/internal/middleware"
)

func Router(db *db.DB, tpl *template.Template, envConfig config.EnvConfig) *mux.Router {

	r := mux.NewRouter()

	// All users have access but shows different UI based on authentication
	r.Handle("/", middleware.Authenticate(ViewIndex(tpl))).Methods("GET")

	// view post page and serve related partials
	r.Handle("/post/{id}", middleware.Authenticate(ViewPost(tpl, db))).Methods("GET")
	r.Handle("/get-post-comments/{id}", PartialComments(db, tpl))
	r.Handle("/submit-comment/{id}", middleware.Authenticate(FormSubmitComment(db, tpl)))

	r.Handle("/login-handler", LoginHandler(db)).Methods("POST")
	r.Handle("/signup", ViewSignup(tpl)).Methods("GET")

	r.Handle("/login", ViewLogin(tpl)).Methods("GET")
	r.Handle("/sign-handler", bunny.SignHandler(envConfig)).Methods("GET")

	// SERVE REQUESTED PARTIALS
	r.Handle("/latest-images", PartialLatestImgs(db, tpl, 0)).Methods("GET")
	r.Handle("/latest-posts-with-img", PartialLatestPostsWithImg(db, tpl, 0))
	r.Handle("/upload", middleware.Authenticate(ViewUpload(tpl))).Methods("GET")

	r.Handle("/create-post", middleware.Authenticate(bunny.UploadImages(envConfig, FormCreatePost(db)))).Methods("POST")

	r.Handle("/new-user", FormCreateUser(db, tpl)).Methods("POST")
	// Serve static files at /static/
	staticFileDirectory := http.Dir("assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET", "HEAD")
	return r
}
