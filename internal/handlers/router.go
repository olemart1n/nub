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
	r.Handle("/", middleware.WithUserContext(ViewIndex(tpl))).Methods("GET")

	// view POST page and serve related partials
	r.Handle("/post/{id}", middleware.WithUserContext(ViewPost(tpl, db))).Methods("GET")
	r.Handle("/get-post-comments/{id}", PartialComments(db, tpl))
	r.Handle("/submit-comment/{id}", middleware.AuthenticationRequired(FormSubmitComment(db, tpl)))

	// SIGN IN
	r.Handle("/sign-in", AuthSignIn(db, tpl)).Methods("POST")
	r.Handle("/sign-in", ViewSignIn(tpl)).Methods("GET")

	// SIGN OUT
	r.Handle("/sign-out", middleware.WithUserContext(AuthSignOut(db, tpl))).Methods("POST")

	// SIGN UP
	r.Handle("/sign-up", ViewSignUp(tpl)).Methods("GET")
	r.Handle("/sign-up", AuthSignUp(db, tpl)).Methods("POST")

	// FOR BUNNY UPLOAD. NOT IN USE
	r.Handle("/sign-handler", bunny.SignHandler(envConfig)).Methods("GET")

	// SERVE REQUESTED PARTIALS
	// r.Handle("/latest-images", PartialLatestImgs(db, tpl, 0)).Methods("GET")
	r.Handle("/latest-posts-with-img", PartialLatestPostsWithImg(db, tpl, 0))

	// CREATE POST
	r.Handle("/upload", middleware.AuthenticationRequired(ViewUpload(tpl))).Methods("GET")
	r.Handle("/create-post", middleware.AuthenticationRequired(bunny.UploadImages(envConfig, FormCreatePost(db, tpl)))).Methods("POST")

	// Serve static files at /static/
	staticFileDirectory := http.Dir("assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET", "HEAD")
	return r
}
