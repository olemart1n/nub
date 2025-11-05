package db

import (
	"context"
)

func (db *DB) CreateComment(ctx context.Context, userID int, postID int, content string) (CommentWithUser, error) {
	var comment CommentWithUser

	err := db.Pool.QueryRow(ctx, `
        WITH inserted AS (
            INSERT INTO comments (post_id, user_id, content)
            VALUES ($1, $2, $3)
            RETURNING id, post_id, user_id, content, created_at
        )
        SELECT 
            i.id, i.post_id, i.user_id, i.content, i.created_at, u.username
        FROM inserted i
        LEFT JOIN users u ON i.user_id = u.id
    `, postID, userID, content).Scan(
		&comment.ID,
		&comment.PostID,
		&comment.UserID,
		&comment.Content,
		&comment.CreatedAt,
		&comment.Username,
	)

	return comment, err
}
