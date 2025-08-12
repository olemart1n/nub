package db

import (
	"context"
	"fmt"
	"strings"
)

func (db *DB) CreatePost(ctx context.Context, userID int, title string, location string, tags []string, imageURLs []string) error {
	// START A TRANSACTION. ALL INSERT MUST SUCCEED.
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		fmt.Println("could not begin a database transaction")
		return err
	}
	defer tx.Rollback(ctx)

	// Insert post
	var postID int
	err = tx.QueryRow(ctx,
		"INSERT INTO posts (user_id, title, location) VALUES ($1, $2, $3) RETURNING id",
		userID, title, location).Scan(&postID)
	if err != nil {
		fmt.Println("could not insert into posts")
		return err
	}

	// ---------------------
	// BULK INSERT imageURLs
	// ---------------------
	imageValues := []any{}
	imageUrlsPlaceholders := []string{}
	i := 1
	for _, url := range imageURLs {
		url = strings.TrimSpace(url)
		if url == "" {
			continue
		}
		imageValues = append(imageValues, postID, url)
		imageUrlsPlaceholders = append(imageUrlsPlaceholders, fmt.Sprintf("($%d, $%d)", i, i+1))
		i += 2
	}

	if len(imageValues) > 0 {
		_, err = tx.Exec(ctx, fmt.Sprintf(
			"INSERT INTO images (post_id, image_url) VALUES %s",
			strings.Join(imageUrlsPlaceholders, ", ")), imageValues...)
		if err != nil {
			return err
		}
	}

	// ------------------
	// BULK INSERT tags
	// ------------------

	tagValues := []any{}
	tagPlaceholders := []string{}
	i = 1
	cleanTags := []string{}

	for _, t := range tags {
		t = strings.TrimSpace(strings.ToLower(t))
		if t == "" {
			continue
		}
		cleanTags = append(cleanTags, t)
		tagValues = append(tagValues, t)
		tagPlaceholders = append(tagPlaceholders, fmt.Sprintf("($%d)", i))
		i++
	}

	if len(cleanTags) > 0 {
		// Insert all tags, ignore duplicates
		_, err = tx.Exec(ctx,
			fmt.Sprintf("INSERT INTO tags (name) VALUES %s ON CONFLICT (name) DO NOTHING",
				strings.Join(tagPlaceholders, ", ")),
			tagValues...)
		if err != nil {
			fmt.Println("could not insert into tags")
			return err
		}

		// Get IDs for all tags (both existing and newly inserted)
		tagIDs := []int{}
		query := fmt.Sprintf("SELECT id FROM tags WHERE name IN (%s)",
			strings.Join(makePlaceholders(len(cleanTags)), ", "))
		rows, err := tx.Query(ctx, query, stringSliceToInterfaces(cleanTags)...)
		if err != nil {
			fmt.Println("could not select id from tags")
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			if err := rows.Scan(&id); err != nil {
				return err
			}
			tagIDs = append(tagIDs, id)
		}

		// Link tags to post
		linkValues := []any{}
		linkPlaceholders := []string{}
		i = 1
		for _, tagID := range tagIDs {
			linkValues = append(linkValues, postID, tagID)
			linkPlaceholders = append(linkPlaceholders, fmt.Sprintf("($%d, $%d)", i, i+1))
			i += 2
		}

		if len(linkValues) > 0 {
			_, err = tx.Exec(ctx,
				fmt.Sprintf("INSERT INTO post_tags (post_id, tag_id) VALUES %s ON CONFLICT DO NOTHING",
					strings.Join(linkPlaceholders, ", ")),
				linkValues...)
			if err != nil {
				fmt.Println("could not insert into post_tags")
				return err
			}
		}
	}

	return nil
}

// Helper to make placeholders like $1, $2, $3...
func makePlaceholders(n int) []string {
	out := make([]string, n)
	for i := 1; i <= n; i++ {
		out[i-1] = fmt.Sprintf("$%d", i)
	}
	return out
}

// Convert []string to []interface{}
func stringSliceToInterfaces(slice []string) []any {
	out := make([]any, len(slice))
	for i, v := range slice {
		out[i] = v
	}
	return out
}
