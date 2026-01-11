package db

import (
	"context"
	"fmt"
)

func (db *DB) GetImageWithDelay(ctx context.Context) (string, error) {

	var imageURL string
sql := `SELECT  image_url FROM images WHERE id = 17`
	row := db.Pool.QueryRow(ctx, sql)

	
	err := row.Scan(&imageURL)
	if err != nil {
		fmt.Print("error scanning row: ", err)
		return "Error occured when fetching imageurl", err
	}
	return imageURL, nil
}
