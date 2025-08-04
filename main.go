package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/olemart1n/nub/router"
)

func main() {
	var pool *pgxpool.Pool
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	r := router.NewRouter(pool, tmpl)

	log.Print("Listening on port 8080")
	http.ListenAndServe(":8080", r)
}

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
