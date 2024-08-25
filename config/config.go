package config

import "github.com/codebyaadi/rss-feed-agg/internal/database"

type ApiConfig struct {
	DB *database.Queries
}