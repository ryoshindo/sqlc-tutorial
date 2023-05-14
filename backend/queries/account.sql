-- name: FindById :one
SELECT * FROM accounts WHERE id = $1 LIMIT 1;

-- name: FindByEmail :one
SELECT * FROM accounts WHERE email = $1 LIMIT 1;

-- name: CreateAccount :exec
INSERT INTO accounts (id, email, created_at, updated_at) VALUES ($1, $2, $3, $4);
