package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/olemart1n/nub/config"
	"github.com/olemart1n/nub/internal/db"
	"github.com/olemart1n/nub/internal/handlers"
	"github.com/olemart1n/nub/internal/session"
)

// var views = jet.NewSet(
// 	jet.NewOSFileSystemLoader("./views"),
// 	jet.InDevelopmentMode(), // Fjern for produksjon
// )

func main() {
	envConfig := config.LoadEnvConfig()
	database, err := db.Connect(envConfig.DatabaseURL)
	if err != nil {
		log.Fatal("Error when connecting to database", err)
	}

	session.InitRedis("localhost:6379")

	tmpl := template.Must(template.ParseGlob("templates/**/*.html"))
	r := handlers.Router(database, tmpl, envConfig)

	log.Print("Listening on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server error: ", err)
	}
}
