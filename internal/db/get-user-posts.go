package db

import "context"

func (db *DB) GetUserPosts(username string) ([]Post, error) {
	ctx := context.Background()

	sql := `
        SELECT p.id, p.user_id, p.title, p.location, p.created_at,
               ARRAY(SELECT t.name
                     FROM tags t
                     JOIN post_tags pt ON pt.tag_id = t.id
                     WHERE pt.post_id = p.id) AS tags
        FROM posts p
        JOIN users u ON p.user_id = u.id
        WHERE u.username = $1
        ORDER BY p.created_at DESC
    `

	rows, err := db.Pool.Query(ctx, sql, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
