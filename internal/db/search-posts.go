package db

import "context"

func (db *DB) SearchPosts(query string) ([]Post, error) {
	ctx := context.Background()

	sql := `
        SELECT p.id, p.user_id, p.title, p.location, p.image_url, p.created_at,
               ARRAY(SELECT t.name
                     FROM tags t
                     JOIN post_tags pt ON pt.tag_id = t.id
                     WHERE pt.post_id = p.id) AS tags
        FROM posts p
        WHERE to_tsvector('english', p.title || ' ' || COALESCE(p.location, '')) @@ plainto_tsquery($1)
        ORDER BY p.created_at DESC
    `

	rows, err := db.Pool.Query(ctx, sql, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []Post
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Location, &p.ImageURL, &p.CreatedAt, &p.Tags)
		if err != nil {
			return nil, err
		}
		results = append(results, p)
	}
	return results, nil
}
