package db

import (
	"context"
	"time"
)

func (db *DB) SearchPostsWithImg(ctx context.Context, query string, page int) ([]PostWithImg, error) {
	const pageSize = 12
	offset := page * pageSize
	sql := `
        SELECT 
            p.id,
            p.title,
            p.location,
            p.created_at,
            i.image_url,
            i.created_at AS image_created_at,
            img_counts.total_images
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
        WHERE to_tsvector('english', p.title || ' ' || COALESCE(p.location, ''))
              @@ plainto_tsquery($1)
        ORDER BY p.created_at DESC
				LIMIT $2 OFFSET $3
    `

	rows, err := db.Pool.Query(ctx, sql, query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []PostWithImg
	for rows.Next() {
		var p PostWithImg
		var createdAt time.Time
		err := rows.Scan(
			&p.Post.ID,
			&p.Post.Title,
			&p.Post.Location,
			&p.Post.CreatedAt,
			&p.Image.ImageURL,
			&createdAt,
			&p.ImageCount,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, p)
	}

	return results, nil
}
