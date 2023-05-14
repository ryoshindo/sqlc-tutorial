package repository

import (
	"context"

	"github.com/ryoshindo/sqlc-tutorial/backend/model"
)

type SessionRepository interface {
	FindByToken(ctx context.Context, token string) (*model.Session, error)
	Create(ctx context.Context, session *model.Session) error
	Delete(ctx context.Context, id string) error
}
