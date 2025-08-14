// Package handlers handles http requests
package handlers

import (
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/mux"
	"github.com/olemart1n/nub/config"
	"github.com/olemart1n/nub/internal/db"
	"github.com/olemart1n/nub/internal/handlers/bunny"
	"github.com/olemart1n/nub/internal/middleware"
)

func Router(db *db.DB, views *jet.Set, envConfig config.EnvConfig) *mux.Router {

	r := mux.NewRouter()

	// All users have access but shows different UI based on authentication
	r.Handle("/", middleware.Authenticate(ViewIndex(views))).Methods("GET")
	r.Handle("/post/{id}", middleware.Authenticate(ViewPost(views, db))).Methods("GET")

	r.Handle("/login-handler", LoginHandler(db)).Methods("POST")
	r.Handle("/signup", ViewSignup(views)).Methods("GET")

	r.Handle("/login", ViewLogin(views)).Methods("GET")
	r.Handle("/sign-handler", bunny.SignHandler(envConfig)).Methods("GET")

	r.Handle("/latest-images", PartialLatestImgs(db, views, 0)).Methods("GET")

	r.Handle("/upload", middleware.Authenticate(ViewUpload(views))).Methods("GET")

	r.Handle("/create-post", middleware.Authenticate(bunny.UploadImages(envConfig, FormCreatePost(db)))).Methods("POST")

	r.Handle("/new-user", HxCreateUser(db, views)).Methods("POST")
	// Serve static files at /static/
	staticFileDirectory := http.Dir("assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET", "HEAD")
	return r
}
