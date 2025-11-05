package db

import (
	"context"
	"fmt"
)

func (db *DB) GetPost(ctx context.Context, id int) (Post, error) {

	sql := `SELECT p.id, p.user_id, p.title, p.location, p.created_at, u.username FROM posts p JOIN users u ON u.id = p.user_id WHERE p.id = $1`

	row := db.Pool.QueryRow(ctx, sql, id)

	var result Post
	err := row.Scan(&result.ID, &result.UserID, &result.Title, &result.Location, &result.CreatedAt, &result.Username)
	if err != nil {
		fmt.Print("error scanning row: ", err)
		return Post{}, err
	}
	return result, nil
}
