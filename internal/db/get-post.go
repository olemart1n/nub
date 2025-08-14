package db

import (
	"context"
	"fmt"
)

func (db *DB) GetPost(ctx context.Context, id int) (Post, error) {

	sql := `SELECT id, user_id, title, location, created_at FROM posts WHERE id = $1`

	row := db.Pool.QueryRow(ctx, sql, id)

	var result Post
	err := row.Scan(&result.ID, &result.UserID, &result.Title, &result.Location, &result.CreatedAt)
	if err != nil {
		fmt.Print("error scanning row: ", err)
		return Post{}, err
	}
	return result, nil
}
