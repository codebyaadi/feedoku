package handlers

import (
	"time"

	"github.com/codebyaadi/rss-feed-agg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey    string    `json:"api_key"`
}

// convertDatabaseUserToAPIUser converts a database user model to an API user model.
//
// Parameters:
//   - dbUser: The user model from the database, typically containing fields like ID, Name, CreatedAt, and UpdatedAt.
//
// Returns:
//   - A User struct containing the same ID, Name, CreatedAt, and UpdatedAt fields, formatted for API responses.
//
// This function maps the fields from the database-specific User struct to the API-specific
// User struct, ensuring the data is properly formatted for JSON serialization and API responses.
func convertDatabaseUserToAPIUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

// convertDatabaseFeedToAPIFeed converts a database feed model to an API feed model.
//
// Parameters:
//   - dbFeed: The feed model from the database, typically containing fields like ID, Name, CreatedAt, UpdatedAt, Url, and UserID.
//
// Returns:
//   - A Feed struct containing the same ID, Name, CreatedAt, UpdatedAt, Url, and UserID fields, formatted for API responses.
//
// This function maps the fields from the database-specific Feed struct to the API-specific
// Feed struct, ensuring the data is properly formatted for JSON serialization and API responses.
func convertDatabaseFeedToAPIFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		Name:      dbFeed.Name,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

// convertDatabaseFeedsToAPIFeeds converts a slice of database feed models to a slice of API feed models.
//
// Parameters:
//   - dbFeeds: A slice of database feed models, each containing fields like ID, Name, CreatedAt, UpdatedAt, Url, and UserID.
//
// Returns:
//   - A slice of Feed structs, each containing the same ID, Name, CreatedAt, UpdatedAt, Url, and UserID fields,
//     formatted for API responses.
//
// This function iterates over each database feed model, converts it to an API feed model using
// convertDatabaseFeedToAPIFeed, and returns a slice of these API feed models.
func convertDatabaseFeedsToAPIFeeds(dbFeeds []database.Feed) []Feed {
	feeds := make([]Feed, len(dbFeeds))
	for i, dbFeed := range dbFeeds {
		feeds[i] = convertDatabaseFeedToAPIFeed(dbFeed)
	}
	return feeds
}

// func convertDatabaseFeedsToAPIFeeds(dbFeeds []database.Feed) []Feed {
// 	feeds := []Feed{}
// 	for _, dbFeed := range dbFeeds {
// 		feeds = append(feeds, convertDatabaseFeedToAPIFeed(dbFeed))
// 	}
// 	// Suggested code may be subject to a license. Learn more: ~LicenseLog:1448372978.
// 	return feeds
// }
