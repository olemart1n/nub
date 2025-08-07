package db

import (
	"context"
	"strings"
)

func (db *DB) CreatePost(userID int, title, location, imageURL string, tags []string) error {
	ctx := context.Background()

	// Insert post
	var postID int
	err := db.Pool.QueryRow(ctx,
		"INSERT INTO posts (user_id, title, location, image_url) VALUES ($1, $2, $3, $4) RETURNING id",
		userID, title, location, imageURL).Scan(&postID)
	if err != nil {
		return err
	}

	// Insert tags & link them
	for _, t := range tags {
		t = strings.TrimSpace(strings.ToLower(t))
		if t == "" {
			continue
		}

		var tagID int
		err := db.Pool.QueryRow(ctx,
			"INSERT INTO tags (name) VALUES ($1) ON CONFLICT (name) DO UPDATE SET name=EXCLUDED.name RETURNING id",
			t).Scan(&tagID)
		if err != nil {
			return err
		}

		_, err = db.Pool.Exec(ctx,
			"INSERT INTO post_tags (post_id, tag_id) VALUES ($1, $2) ON CONFLICT DO NOTHING",
			postID, tagID)
		if err != nil {
			return err
		}
	}

	return nil
}
