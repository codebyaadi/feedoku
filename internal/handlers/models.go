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
	}
}
