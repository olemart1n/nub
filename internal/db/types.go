// Package db handles database logic
package db

import (
	"database/sql"
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
	Email     sql.NullString
}

type Post struct {
	ID        int
	UserID    int
	Title     string
	Location  string
	CreatedAt time.Time
	Username  string
}

type Image struct {
	ID        int
	PostID    int
	ImageURL  string
	CreatedAt time.Time
}

type Comment struct {
	ID        int
	PostID    int
	UserID    int
	Content   string
	CreatedAt time.Time
}
type CommentWithUser struct {
	Comment
	Username *string
}
type PostWithImg struct {
	Post       Post
	Image      Image
	ImageCount int
}
