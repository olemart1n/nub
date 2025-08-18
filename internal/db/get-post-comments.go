package db

import (
	"context"
	"fmt"
)

func (db *DB) GetPostComments(ctx context.Context, postID int) ([]CommentWithUser, error) {

	var comments []CommentWithUser

	sql := `	SELECT 
  comments.id,
  comments.user_id,
  comments.content,
  comments.created_at,
  users.username
FROM comments
JOIN users ON comments.user_id = users.id
WHERE comments.post_id = $1
ORDER BY comments.created_at DESC`

	rows, err := db.Pool.Query(ctx, sql, postID)
	if err != nil {
		return comments, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment CommentWithUser
		err := rows.Scan(&comment.ID, &comment.UserID, &comment.Content, &comment.CreatedAt, &comment.Username)
		if err != nil {
			fmt.Print("error scanning row: ", err)
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil

}
