package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
)

func insertImages(ctx context.Context, tx pgx.Tx, postID int, imageURLs []string) error {
	var values []any
	var placeholders []string
	i := 1

	for _, url := range imageURLs {
		url = strings.TrimSpace(url)
		if url == "" {
			continue
		}
		values = append(values, postID, url)
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d)", i, i+1))
		i += 2
	}

	if len(values) == 0 {
		return nil
	}

	_, err := tx.Exec(ctx, fmt.Sprintf(
		"INSERT INTO images (post_id, image_url) VALUES %s",
		strings.Join(placeholders, ", ")), values...)
	return err
}
