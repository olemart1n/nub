package db

import (
	"context"
	"fmt"
	"time"
)

func (db *DB) GetLatestPostsWithImg(ctx context.Context, page int) ([]PostWithImg, error) {
	const pageSize = 12
	offset := page * pageSize

	query := `
	SELECT p.title, p.created_at, p.id, p.location, i.image_url, i.created_at, img_counts.total_images 
	FROM posts p 
	JOIN LATERAL (
	    SELECT id, image_url, created_at
	    FROM images 
	    WHERE post_id = p.id 
	    ORDER BY created_at ASC 
	    LIMIT 1
	) i ON TRUE
	JOIN (
	    SELECT post_id, COUNT(*) AS total_images 
	    FROM images 
	    GROUP BY post_id
	) img_counts ON img_counts.post_id = p.id
	ORDER BY p.created_at DESC 
	LIMIT $1 OFFSET $2`

	rows, err := db.Pool.Query(ctx, query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []PostWithImg

	for rows.Next() {
		var P PostWithImg
		var createdAt time.Time

		err := rows.Scan(
			&P.Post.Title,
			&P.Post.CreatedAt,
			&P.Post.ID,
			&P.Post.Location,
			&P.Image.ImageURL,
			&createdAt,
			&P.ImageCount,
		)
		if err != nil {
			return nil, err
		}

		// Append version param based on created_at
		P.Image.ImageURL = fmt.Sprintf("%s?v=%d", P.Image.ImageURL, createdAt.Unix())

		results = append(results, P)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
