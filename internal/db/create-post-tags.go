package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
)

func insertTags(ctx context.Context, tx pgx.Tx, postID int, tags []string) error {
	var cleanTags []string
	var values []any
	var placeholders []string
	i := 1

	for _, t := range tags {
		t = strings.TrimSpace(strings.ToLower(t))
		if t == "" {
			continue
		}
		cleanTags = append(cleanTags, t)
		values = append(values, t)
		placeholders = append(placeholders, fmt.Sprintf("($%d)", i))
		i++
	}

	if len(cleanTags) == 0 {
		return nil
	}

	// Insert all tags (ignore duplicates)
	_, err := tx.Exec(ctx,
		fmt.Sprintf("INSERT INTO tags (name) VALUES %s ON CONFLICT (name) DO NOTHING",
			strings.Join(placeholders, ", ")),
		values...)
	if err != nil {
		fmt.Println("could not insert into tags")
		return err
	}

	// Get IDs for tags
	tagIDs, err := fetchTagIDs(ctx, tx, cleanTags)
	if err != nil {
		return err
	}

	// Link tags to post
	return linkTagsToPost(ctx, tx, postID, tagIDs)
}

func fetchTagIDs(ctx context.Context, tx pgx.Tx, tags []string) ([]int, error) {
	query := fmt.Sprintf("SELECT id FROM tags WHERE name IN (%s)",
		strings.Join(makePlaceholders(len(tags)), ", "))
	rows, err := tx.Query(ctx, query, stringSliceToInterfaces(tags)...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func linkTagsToPost(ctx context.Context, tx pgx.Tx, postID int, tagIDs []int) error {
	if len(tagIDs) == 0 {
		return nil
	}

	var values []any
	var placeholders []string
	i := 1
	for _, tagID := range tagIDs {
		values = append(values, postID, tagID)
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d)", i, i+1))
		i += 2
	}

	_, err := tx.Exec(ctx,
		fmt.Sprintf("INSERT INTO post_tags (post_id, tag_id) VALUES %s ON CONFLICT DO NOTHING",
			strings.Join(placeholders, ", ")),
		values...)
	return err
}
