package db

import (
	"context"
	"fmt"
)

func (db *DB) CreatePost(ctx context.Context, userID int, title string, location string, tags []string, imageURLs []string) (int, error) {
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		fmt.Println("could not begin a database transaction")
		return 0, err
	}
	defer tx.Rollback(ctx)

	// Insert post
	var postID int
	err = tx.QueryRow(ctx,
		"INSERT INTO posts (user_id, title, location) VALUES ($1, $2, $3) RETURNING id",
		userID, title, location).Scan(&postID)
	if err != nil {
		fmt.Println("could not insert into posts")
		return 0, err
	}

	// Insert images
	if err := insertImages(ctx, tx, postID, imageURLs); err != nil {
		return 0, err
	}

	// Insert tags
	if err := insertTags(ctx, tx, postID, tags); err != nil {
		return 0, err
	}

	return postID, tx.Commit(ctx)
}
