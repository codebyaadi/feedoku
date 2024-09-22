package auth

import (
	"context"
	"time"

	"github.com/codebyaadi/rss-feed-agg/internal/redis"
	"github.com/google/uuid"
)

var ctx = context.Background()

func SetRefreshToken(userId uuid.UUID, refreshToken string, expiration time.Duration) error {
	return redis.Client.Set(ctx, userKey(userId), refreshToken, expiration).Err()
}

func GetRefreshToken(userId uuid.UUID) (string, error) {
	return redis.Client.Get(ctx, userKey(userId)).Result()
}

func DeleteRefreshToken(userId uuid.UUID) error {
	return redis.Client.Del(ctx, userKey(userId)).Err()
}

func userKey(userId uuid.UUID) string {
	return "user:refresh_token:" + userId.String()
}
