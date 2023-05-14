package db

import (
	"context"
	"database/sql"
)

type contextKey struct {
	name string
}

func Set(ctx context.Context, db *sql.DB) context.Context {
	return context.WithValue(ctx, contextKey{"db"}, db)
}

func Get(ctx context.Context) *sql.DB {
	db, ok := ctx.Value(contextKey{"db"}).(*sql.DB)
	if !ok {
		panic("db connection is not found")
	}
	return db
}
