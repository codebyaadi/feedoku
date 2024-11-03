package models

import (
	"time"

	"github.com/codebyaadi/rss-feed-agg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ApiKey       string    `json:"api_key"`
	Email        string    `json:"email"`
	AccessToken  string    `json:"access_token,omitempty"`
	RefreshToken string    `json:"refresh_token,omitempty"`
}

func ConvertDatabaseUserToAPIUser(dbUser database.User, accessToken, refreshToken string) User {
	user := User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey:    dbUser.ApiKey,
		Email:     dbUser.Email,
	}

	if accessToken != "" {
		user.AccessToken = accessToken
	}
	if refreshToken != "" {
		user.RefreshToken = refreshToken
	}

	return user
}
