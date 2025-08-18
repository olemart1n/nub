package db

import (
	"context"
)

func (db *DB) GetLatestPostsWithImg(ctx context.Context, page int) ([]PostWithImg, error) {

	const pageSize = 12
	offset := page * pageSize

	query := `
	SELECT p.title, p.id, p.location, i.image_url, img_counts.total_images 
	FROM posts p JOIN LATERAL (SELECT id, image_url FROM images WHERE post_id = p.id ORDER BY created_at ASC LIMIT 1) i ON TRUE
	JOIN (SELECT post_id, COUNT(*) AS total_images FROM images GROUP BY post_id) img_counts ON img_counts.post_id = p.id
	ORDER BY p.created_at DESC LIMIT $1 OFFSET $2`
	rows, err := db.Pool.Query(ctx, query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []PostWithImg

	for rows.Next() {
		var P PostWithImg
		err := rows.Scan(&P.Post.Title, &P.Post.ID, &P.Post.Location, &P.Image.ImageURL, &P.ImageCount)
		if err != nil {
			return nil, err
		}
		results = append(results, P)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
