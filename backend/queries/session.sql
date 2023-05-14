-- name: FindByToken :one
SELECT * FROM sessions WHERE token = $1;

-- name: CreateSession :exec
INSERT INTO sessions (id, account_id, token, expires_at, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6);

-- name: DeleteSession :exec
DELETE FROM sessions WHERE id = $1;
