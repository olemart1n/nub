// Package db handles database logic
package db

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

type User struct {
	ID        int
	Username  string
	CreatedAt time.Time
}

type Post struct {
	ID        int
	UserID    int
	Title     string
	Location  string
	ImageURL  string
	CreatedAt time.Time
	Tags      []string
}
