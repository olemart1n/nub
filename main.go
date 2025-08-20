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

func main() {
	envConfig := config.LoadEnvConfig()
	database, err := db.Connect(envConfig.DatabaseURL)
	if err != nil {
		log.Fatal("Error when connecting to database", err)
	}

	session.InitRedis(envConfig.RedisURL)

	tmpl := template.Must(template.ParseGlob("templates/**/*.html"))
	r := handlers.Router(database, tmpl, envConfig)

	log.Print("Listening on port: " + envConfig.PORT)
	if err := http.ListenAndServe(":"+envConfig.PORT, r); err != nil {
		log.Fatal("Server error: ", err)
	}
}
