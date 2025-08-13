package db

import (
	"context"
	"fmt"
)

func (db *DB) GetLatestImages(ctx context.Context, page int) ([]Image, error) {

	const pageSize = 12
	offset := page * pageSize

	query := `SELECT image_url, post_id FROM images ORDER BY id DESC LIMIT $1 OFFSET $2`

	rows, err := db.Pool.Query(ctx, query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []Image

	for rows.Next() {
		var i Image
		err := rows.Scan(&i.ImageURL, &i.PostID)
		if err != nil {
			fmt.Print("some error happened")
			return nil, err
		}
		fmt.Println("Fetched image:", i.ImageURL)
		results = append(results, i)
	}
	return results, nil
}
