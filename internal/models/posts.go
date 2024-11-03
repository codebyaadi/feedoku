package models

import (
	"time"

	"github.com/codebyaadi/rss-feed-agg/internal/database"
	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ConvertDatabasePostToAPIPost(dbPost database.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return Post{
		ID:          dbPost.ID,
		Title:       dbPost.Title,
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID,
		PublishedAt: dbPost.PublishedAt,
		Description: description,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
	}
}

func ConvertDatabasePostsToAPIPosts(dbPosts []database.Post) []Post {
	posts := make([]Post, len(dbPosts))
	for i, dbPost := range dbPosts {
		posts[i] = ConvertDatabasePostToAPIPost(dbPost)
	}
	return posts
}
