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
	CreatedAt time.Time
}

type Image struct {
	ID        int
	PostID    int
	ImageURL  string
	CreatedAt time.Time
}

type Comment struct {
	id        int
	postID    int
	userID    int
	content   string
	createdAt time.Time
}
