package db

import (
	"context"

	generatedsqlc "github.com/ryoshindo/sqlc-tutorial/backend/model/sqlc"
)

type contextKey struct {
	name string
}

func Set(ctx context.Context, q *generatedsqlc.Queries) context.Context {
	return context.WithValue(ctx, contextKey{"db"}, q)
}

func Get(ctx context.Context) *generatedsqlc.Queries {
	q, ok := ctx.Value(contextKey{"db"}).(*generatedsqlc.Queries)
	if !ok {
		panic("db connection is not found")
	}
	return q
}
