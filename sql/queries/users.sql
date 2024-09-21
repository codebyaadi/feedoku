-- name: CreateUser :one
INSERT INTO users (id, name, email, password_hash, created_at, updated_at, oauth_provider, oauth_id, api_key)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8,
    encode(sha256(random()::text::bytea), 'hex')
)
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM users
WHERE api_key = $1 LIMIT 1;

-- name: UpdateOAuthID :one
UPDATE users
SET oauth_provider = $1, oauth_id = $2
WHERE email = $3
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUserPassword :one
UPDATE users
SET password_hash = $1
WHERE email = $2
RETURNING *;