-- name: FindById :one
SELECT * FROM accounts WHERE id = $1 LIMIT 1;

-- name: FindByEmail :one
SELECT * FROM accounts WHERE email = $1 LIMIT 1;
