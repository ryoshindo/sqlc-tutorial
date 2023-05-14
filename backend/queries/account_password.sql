-- name: FindPasswordsByAccountId :many
SELECT * FROM account_passwords WHERE account_id = $1;
