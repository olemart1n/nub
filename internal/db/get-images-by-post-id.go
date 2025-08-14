package db

import (
	"context"
	"fmt"
)

func (db *DB) GetImagesByPostID(ctx context.Context, postID int) ([]Image, error) {
	sql := `SELECT id, image_url, post_id, created_at FROM images WHERE post_id = $1 ORDER BY created_at DESC`

	rows, err := db.Pool.Query(ctx, sql, postID)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var images []Image
	for rows.Next() {
		var img Image
		err := rows.Scan(&img.ID, &img.ImageURL, &img.PostID, &img.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		images = append(images, img)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return images, nil
}
