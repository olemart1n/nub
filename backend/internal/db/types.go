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
	ID        int            `json:"id"`
	Username  string         `json:"username"`
	CreatedAt time.Time      `json:"createdAt"`
	Email     sql.NullString `json:"email,omitempty"`
}

type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userId"`
	Title     string    `json:"title"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"createdAt"`
	Username  string    `json:"username"`
}

type Image struct {
	ID        int       `json:"id"`
	PostID    int       `json:"postId"`
	Country   string    `json:"country"`
	ImageURL  string    `json:"imageUrl"`
	CreatedAt time.Time `json:"createdAt"`
}

type Comment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"postId"`
	UserID    int       `json:"userId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type CommentWithUser struct {
	Comment
	Username *string `json:"username,omitempty"`
}

type PostWithImg struct {
	Post       Post  `json:"post"`
	Image      Image `json:"image"`
	ImageCount int   `json:"imageCount"`
}
