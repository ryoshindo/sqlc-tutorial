// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: account_password.sql

package generatedsqlc

import (
	"context"
)

const findPasswordsByAccountId = `-- name: FindPasswordsByAccountId :many
SELECT id, account_id, hash, salt, created_at, updated_at FROM account_passwords WHERE account_id = $1
`

func (q *Queries) FindPasswordsByAccountId(ctx context.Context, accountID string) ([]AccountPassword, error) {
	rows, err := q.db.QueryContext(ctx, findPasswordsByAccountId, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AccountPassword
	for rows.Next() {
		var i AccountPassword
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Hash,
			&i.Salt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
