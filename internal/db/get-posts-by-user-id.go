package db

import (
	"context"
)

func (db *DB) GetPostsByUserID(ctx context.Context, userID int) ([]Post, error) {

	sql := `SELECT id, user_id, title, location, created_at FROM posts WHERE user_id = $1`

	rows, err := db.Pool.Query(ctx, sql, userID)

	if err != nil {
		return nil, err
	}
	var results []Post

	for rows.Next() {
		var p Post
		err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Location, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, p)
	}
	return results, nil
}
